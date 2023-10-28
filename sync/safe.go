package sync

import (
	"fmt"
	"time"
)

// ------------------- SafeWrapperSlice ------------------------
// Struct for holding and guarding any type of elements in slice.
type SafeWrapperSlice[SharedDataType any] struct {
	data      []SharedDataType
	dataGuard *Waiter
}

func NewSafeWrapperSlice[SharedDataType any]() *SafeWrapperSlice[SharedDataType] {
	return &SafeWrapperSlice[SharedDataType]{
		dataGuard: NewWaiter(func() bool { return true }),
	}
}

func (sw *SafeWrapperSlice[SharedDataType]) AppendNewElement(newElement SharedDataType) {
	sw.dataGuard.Lock()
	defer sw.dataGuard.Unlock()

	sw.data = append(sw.data, newElement)
	sw.dataGuard.Broadcast()
}

func (sw *SafeWrapperSlice[SharedDataType]) AppendMany(newElements *[]SharedDataType) {
	sw.dataGuard.Lock()
	defer sw.dataGuard.Unlock()

	sw.data = append(sw.data, (*newElements)...)
	sw.dataGuard.Broadcast()
}

func (sw *SafeWrapperSlice[SharedDataType]) WaitAndGetElement(desiredElementIndex int) SharedDataType {
	//TODO: Need additional checks for index { if desiredElementIndex < 0 }.
	sw.dataGuard.Lock()
	defer sw.dataGuard.Unlock()
	sw.dataGuard.WaitForCurrentCondition(func() bool {
		return desiredElementIndex < len(sw.data)
	})

	return sw.data[desiredElementIndex]
}

func (sw *SafeWrapperSlice[SharedDataType]) TryGetElement(desiredElementIndex int, element *SharedDataType) {
	sw.dataGuard.Lock()
	defer sw.dataGuard.Unlock()
	if desiredElementIndex < 0 || desiredElementIndex >= len(sw.data) {
		return
	}
	d := sw.data[desiredElementIndex]
	element = &d
}

func (sw *SafeWrapperSlice[SharedDataType]) WaitAndGetRangeOfElements(from, to int) []SharedDataType {
	//TODO: Before mutex lock and unlock need additional checks for index { if from > to || from < 0 || to < 0 }.
	sw.dataGuard.Lock()
	defer sw.dataGuard.Unlock()

	// terminal condition.
	if from > to || from < 0 || to < 0 {
		return nil
	}

	sw.dataGuard.WaitForCurrentCondition(func() bool {
		return to <= len(sw.data) && len(sw.data) != 0
	})

	var desiredElements []SharedDataType

	if from == to {
		desiredElements = append(desiredElements, sw.data[from])
	} else {
		desiredElements = append(desiredElements, sw.data[from:to]...)
	}
	return desiredElements
}

func (sw *SafeWrapperSlice[SharedDataType]) Size() int {
	sw.dataGuard.Lock()
	defer sw.dataGuard.Unlock()
	return len(sw.data)
}

func (sw *SafeWrapperSlice[SharedDataType]) Remove(from, to int) error {
	sw.dataGuard.Lock()
	defer sw.dataGuard.Unlock()
	if from > to || from < 0 || to >= len(sw.data) || to < 0 {
		return fmt.Errorf("Index out of range. Indexes %d %d. Length of slice %d", from, to, len(sw.data))
	}
	sw.data = append(sw.data[0:from], sw.data[to+1:len(sw.data)]...)
	return nil
}

func (sw *SafeWrapperSlice[SharedDataType]) Cleanup() {
	sw.dataGuard.Lock()
	defer sw.dataGuard.Unlock()
	sw.data = nil
}

// ----------------------- SafeWrapperMap -------------------
// Struct for holding and guarding any type of value with mapping on comparable type.
type SafeWrapperMap[KeyType comparable, ValueType any] struct {
	data      map[KeyType]ValueType
	dataGuard *Waiter
}

func NewSafeWrapperMap[KeyType comparable, ValueType any]() *SafeWrapperMap[KeyType, ValueType] {
	return &SafeWrapperMap[KeyType, ValueType]{
		data:      make(map[KeyType]ValueType),
		dataGuard: NewWaiter(func() bool { return true }),
	}
}

func (swp *SafeWrapperMap[KeyType, ValueType]) AddNewElement(key KeyType, value ValueType) {
	swp.dataGuard.Lock()
	defer swp.dataGuard.Unlock()

	swp.data[key] = value

	swp.dataGuard.Broadcast()
}

func (swp *SafeWrapperMap[KeyType, ValueType]) AddNewElements(elements map[KeyType]ValueType) {
	swp.dataGuard.Lock()
	defer swp.dataGuard.Unlock()

	for key, value := range elements {
		swp.data[key] = value
	}

	swp.dataGuard.Broadcast()
}

func (swp *SafeWrapperMap[KeyType, ValueType]) TryGetElement(key KeyType) (ValueType, bool) {
	swp.dataGuard.Lock()
	defer swp.dataGuard.Unlock()

	data, ok := swp.data[key]
	return data, ok
}

func (swp *SafeWrapperMap[KeyType, ValueType]) WaitAndGetElement(key KeyType) (ValueType, bool) {
	swp.dataGuard.Lock()
	defer swp.dataGuard.Unlock()

	var data *ValueType
	var ok *bool
	swp.dataGuard.WaitForCurrentCondition(func() bool {
		d, o := swp.data[key]
		data = &d
		ok = &o
		return o
	})
	// if {
	// 	swp.dataGuard.Wait()
	// }
	return *data, *ok
}

func (swp *SafeWrapperMap[KeyType, ValueType]) WaitAndGetElementFor(key KeyType, duration time.Duration) (ValueType, bool) {
	var value ValueType
	var ok bool
	now := time.Now().UnixMilli()
	now += duration.Milliseconds()

	swp.dataGuard.Lock()
	defer swp.dataGuard.Unlock()
	for now > time.Now().UnixMilli() {
		if value, ok = swp.data[key]; ok {
			break
		}
	}
	return value, ok
}

func (swp *SafeWrapperMap[KeyType, ValueType]) TryGetElements(keys []KeyType) map[KeyType]ValueType {
	swp.dataGuard.Lock()
	defer swp.dataGuard.Unlock()

	desired := make(map[KeyType]ValueType)
	for _, key := range keys {
		if value, ok := swp.data[key]; ok {
			desired[key] = value
		}
	}
	return desired
}

func (swp *SafeWrapperMap[KeyType, ValueType]) GetKeys() []KeyType {
	swp.dataGuard.Lock()
	defer swp.dataGuard.Unlock()

	var keys []KeyType
	for k := range swp.data {
		keys = append(keys, k)
	}

	return keys
}

func (swp *SafeWrapperMap[KeyType, ValueType]) Remove(key KeyType) ValueType {
	swp.dataGuard.Lock()
	defer swp.dataGuard.Unlock()

	value := swp.data[key]

	delete(swp.data, key)
	return value
}

// Cleaning all map. Returning removed elements count.
func (swp *SafeWrapperMap[KeyType, ValueType]) Cleanup() int {
	swp.dataGuard.Lock()
	defer swp.dataGuard.Unlock()

	count := 0
	for key := range swp.data {
		delete(swp.data, key)
		count++
	}
	return count
}
