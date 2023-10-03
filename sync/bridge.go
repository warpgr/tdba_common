package sync

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// ----------- SyncBridge ----------------------------------
// Composed sync struct. One controller multiple subscribers.
type SyncBridge[SharedDataType interface{}] struct {
	// Sets true if want to start job. And false when want to finish.
	isInProcess *atomic.Bool
	// Is data taken from channel
	dataTaken *atomic.Bool
	// Before starting need to set delta number,
	// it is a count of goroutines which you want to wait after work finishing.
	// call Done() for finishing current subscribed goroutine.
	syncWork *sync.WaitGroup
	// Module name which running
	name       string
	failReason error
	// passing data through channel.
	sharedChannel chan SharedDataType
	// TODO: Fix this.
	onComplete *func(args ...any)
	onFailure  *func(args ...any)
}

func NewSyncBridge[SharedDataType interface{}](delta int, name string, bufferSize int) *SyncBridge[SharedDataType] {
	sb := SyncBridge[SharedDataType]{
		isInProcess:   &atomic.Bool{},
		dataTaken:     &atomic.Bool{},
		syncWork:      &sync.WaitGroup{},
		name:          name,
		sharedChannel: make(chan SharedDataType, bufferSize),
	}
	sb.isInProcess.Store(true)
	sb.dataTaken.Store(true)

	sb.syncWork.Add(delta)
	return &sb
}

func (sb *SyncBridge[SharedDataType]) Add(delta int) {
	sb.syncWork.Add(delta)
}

func (sb *SyncBridge[SharedDataType]) Wait() {
	sb.syncWork.Wait()
	if sb.onComplete != nil && sb.failReason == nil {
		(*sb.onComplete)()
	} else if sb.failReason != nil && sb.onFailure != nil {
		(*sb.onFailure)()
	}
}

func (sb *SyncBridge[SharedDataType]) Done() {
	sb.syncWork.Done()
}

func (sb *SyncBridge[SharedDataType]) Cancel() {
	sb.isInProcess.Store(false)
}

func (sb *SyncBridge[SharedDataType]) StartWork() {
	sb.isInProcess.Store(true)
}

func (sb *SyncBridge[SharedDataType]) IsInProgress() bool {
	return sb.isInProcess.Load()
}

func (sb *SyncBridge[SharedDataType]) Name() string {
	return sb.name
}

func (sb *SyncBridge[SharedDataType]) Reason() error {
	return sb.failReason
}

func (sb *SyncBridge[SharedDataType]) Failed(err error) {
	sb.failReason = err
}

func (sb *SyncBridge[SharedDataType]) BindCompletionHandler(completionHandler func(args ...any)) {
	sb.onComplete = &completionHandler
}

func (sb *SyncBridge[SharedDataType]) BindFailureHandler(failureHandler func(args ...any)) {
	sb.onFailure = &failureHandler
}

func (sb *SyncBridge[SharedDataType]) RecvData() SharedDataType {
	defer sb.dataTaken.Store(true)
	return <-sb.sharedChannel
}

func (sb *SyncBridge[SharedDataType]) AwaitForData(duration time.Duration) *SharedDataType {
	select {
	case data := <-sb.sharedChannel:
		sb.dataTaken.Store(true)
		return &data
	case <-time.After(duration):
		return nil
	}
}

func (sb *SyncBridge[SharedDataType]) SendData(data SharedDataType) {
	if sb.dataTaken.Load() {
		sb.dataTaken.Store(false)
		sb.sharedChannel <- data
	}
}

// Run modules which cyclic working depends from inProgress flag of bridge.
// And notifying finishing by syncWork.Done().
func RunAsync[SharedDataType any](
	name string,
	f interface{},
	args ...any) *SyncBridge[SharedDataType] {

	bridge := NewSyncBridge[SharedDataType](1, name)
	go func() {
		defer bridge.Done() // for avoiding deadlocks
		switch fn := f.(type) {
		case func(*SyncBridge[SharedDataType], ...any):
			fn(bridge, args...)
		case func(*SyncBridge[SharedDataType]):
			fn(bridge)
		}
	}()

	return bridge
}

func RunAsyncWithBinding[SharedDataType any](nameOfModule string, functionObject *BoundValue) (*SyncBridge[SharedDataType], error) {
	// TODO: make sure types are correct.
	bridge := NewSyncBridge[SharedDataType](1, nameOfModule)
	functionObject.argValues = pushBridgeInArgs[SharedDataType](functionObject.argValues, bridge)
	err := functionObject.CheckArgs()
	if err != nil {
		return nil, err
	}
	bridge.StartWork()
	go func() {
		defer bridge.Done()
		functionObject.Invoke()
	}()

	return bridge, nil
}

func pushBridgeInArgs[SharedDataType any](argValues []reflect.Value, bridge *SyncBridge[SharedDataType]) []reflect.Value {
	var newArgValues []reflect.Value
	bridgeReflected := reflect.ValueOf(bridge)
	newArgValues = append(newArgValues, bridgeReflected)
	return newArgValues
}

func Bind(fn interface{}, args ...interface{}) *BoundValue {
	fnValue := reflect.ValueOf(fn)
	argValues := make([]reflect.Value, len(args))

	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}

	return &BoundValue{
		fnValue:   &fnValue,
		argValues: argValues,
	}
}

// --------------- BoundValue ---------------------------------------
// Struct for binding functions as event handler or deferred running.
type BoundValue struct {
	fnValue   *reflect.Value
	argValues []reflect.Value
}

func (bv *BoundValue) Invoke() {
	bv.fnValue.Call(bv.argValues)
}

func (bv *BoundValue) CheckArgs() error {
	fnType := bv.fnValue.Type()

	if fnType.NumIn() != len(bv.argValues) {
		return fmt.Errorf("Different size of args %d %d", fnType.NumIn(), len(bv.argValues))
	}
	for i := 0; i < fnType.NumIn(); i++ {
		if fnType.In(i).Kind() == (bv.argValues)[i].Type().Kind() {
			return fmt.Errorf("Different type of args %s %s", fnType.In(i).Kind(), (bv.argValues)[i].Type().Kind())
		}
	}
	return nil
}

// ---------- Waiter ----------------------------------------
// Struct for synchronization between goroutines by some condition.
type Waiter struct {
	guard     *sync.Mutex
	isReady   *sync.Cond
	condition func() bool
}

func NewWaiter(condition func() bool) *Waiter {
	w := &Waiter{
		guard:     &sync.Mutex{},
		condition: condition,
	}
	w.isReady = sync.NewCond(w.guard)
	return w
}

func (w *Waiter) WaitForCondition() {
	for !w.condition() {
		w.isReady.Wait()
	}
}

func (w *Waiter) WaitForCurrentCondition(condition func() bool) {
	for !condition() {
		w.isReady.Wait()
	}
}

func (w *Waiter) Wait() {
	w.isReady.Wait()
}

func (w *Waiter) Lock() {
	w.guard.Lock()
}

func (w *Waiter) Unlock() {
	w.guard.Unlock()
}

func (w *Waiter) Signal() {
	w.isReady.Signal()
}

func (w *Waiter) Broadcast() {
	w.isReady.Broadcast()
}
