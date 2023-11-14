package sync

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/warpgr/tdba_common/utils"
)

type SharedStateSuite struct {
	suite.Suite
}

func TestSharedStateSuite(t *testing.T) {
	suite.Run(t, new(SharedStateSuite))
}

func (ss *SharedStateSuite) SetupTest() {
	// Setup any common test fixtures here
}

func (ss *SharedStateSuite) TestAsyncTask() {
	// Your test logic for AsyncTask
}

func (ss *SharedStateSuite) TearDownTest() {
	utils.CleanupLogFiles()
}

func (ss *SharedStateSuite) TestSendAndReceiveWithCancel() {
	ctx := context.Background()
	channelBufSize := 4

	i := 0
	asyncFunc := func(promise Promise[int]) {
		for ; i < channelBufSize+1; i++ {
			promise.Send(i)
		}
		promise.CloseChannel()
		ss.Equal(fmt.Errorf("channel closed"), promise.Send(8))
	}

	future := AsyncTask[int](ctx, channelBufSize, asyncFunc)
	j := 0
	for ; j < channelBufSize+1; j++ {
		ss.Equal(j, future.Get())
	}
	ss.Nil(future.Wait())
	ss.Equal(j, i)

	asyncFunc = func(promise Promise[int]) {
		for ; i < channelBufSize+1; i++ {
			promise.Send(i)
		}
	}
	future = AsyncTask[int](ctx, channelBufSize, asyncFunc)
	j = 0
	for ; j < channelBufSize+1; j++ {
		data, isReceived := future.TryGet(time.Duration(1 * time.Second))
		if isReceived {
			ss.Equal(j, data)
		}
	}
	_, isReceived := future.TryGet(time.Duration(1 * time.Second))
	ss.False(isReceived)
	ss.Nil(future.Wait())
	ss.Equal(j, i)
}

func (ss *SharedStateSuite) TestHandlersWithCancel() {
	ctx := context.Background()
	channelBufSize := 0
	iter := atomic.Int32{}

	expectedError := fmt.Errorf("invalid data")
	future := AsyncTask[int32](ctx, channelBufSize,
		func(promise Promise[int32]) {
			for iter.Load() < 5 {
				promise.OnDone(func() {
					ss.True(iter.CompareAndSwap(iter.Load(), iter.Load()+1))
				})
				promise.Send(iter.Load())
				ss.True(iter.CompareAndSwap(iter.Load(), iter.Load()+1))
			}
		}).
		OnData(
			func(data int32) error {
				if data == 4 {
					return fmt.Errorf("invalid data")
				}
				ss.Equal(iter.Load(), data)
				return nil
			}).
		OnFailure(func(err error) {
			ss.True(iter.CompareAndSwap(iter.Load(), iter.Load()+1))
			ss.Equal(expectedError, err)
		})
	ss.Equal(expectedError, future.Wait())
	ss.Equal(int32(7), iter.Load())

	future = AsyncTask[int32](ctx, channelBufSize,
		func(promise Promise[int32]) {

			promise.OnDone(func() {
				ss.True(iter.CompareAndSwap(iter.Load(), iter.Load()+1))
			})

			for iter.Load() < 10 {
				promise.Send(iter.Load())
				ss.True(iter.CompareAndSwap(iter.Load(), iter.Load()+1))
			}
		}).
		OnData(func(data int32) error {
			ss.Equal(iter.Load(), data)
			return nil
		})

	ss.Nil(future.Wait())
	ss.Equal(int32(11), iter.Load())
}

func (ss *SharedStateSuite) TestHandlersWithTimeout() {
	ctx := context.Background()
	channelBufSize := 0
	cancelAfter := time.Duration(3 * time.Second)

	firstIter := atomic.Int32{}
	firstIter.Store(1)
	future := AsyncTaskWithTimeOut[int32](ctx, channelBufSize, cancelAfter,
		func(promise Promise[int32]) {
			for ; firstIter.Load() < 10; firstIter.Add(1) {
				// Sending just one value 0.
				promise.Send(firstIter.Load())
				// Must canceled during sleeping.
				time.Sleep(10 * time.Second)
			}
		}).
		OnData(func(data int32) error {
			return nil
		})
	ss.Nil(future.Wait())
	ss.Equal(int32(1), firstIter.Load())

	secondIter := atomic.Int32{}
	expectedError := fmt.Errorf("somme error from task")
	future = AsyncTaskWithTimeOut[int32](ctx, channelBufSize, cancelAfter,
		func(promise Promise[int32]) {
			promise.OnDone(func() {
				secondIter.Add(1)
			})
			for ; secondIter.Load() < 10; secondIter.Add(1) {
				promise.Send(int32(secondIter.Load()))
			}

		}).
		OnData(func(data int32) error {
			ss.Equal(secondIter.Load(), data)
			if secondIter.Load() == int32(9) {
				return expectedError
			}
			return nil
		}).
		OnFailure(func(err error) {
			secondIter.Add(1)
			ss.Equal(expectedError, err)
		})

	ss.Equal(expectedError, future.Wait())
	ss.Equal(int32(12), secondIter.Load())
}
