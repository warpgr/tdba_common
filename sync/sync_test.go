package sync

// import (
// 	"fmt"
// 	"sync"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/warpgr/tdba_common/utils"
// )

// func TestSyncBridge(t *testing.T) {
// 	// TODO: Need to test onComplete, onFailure. Test with function binding. And Fixing.
// 	defer utils.CleanupLogFiles()

// 	expectedResultsOfSteps := [6]int{1, 2, 3, 5, 6, 7} // Expected results of steps.
// 	var resultOfSteps []int
// 	iterator := 1

// 	controller := NewSyncBridge[int](1, "test-controller-1", 1)

// 	go func() {
// 		defer controller.Done()
// 		controller.SendData(iterator) // 1
// 		iterator++
// 	}()
// 	resultOfSteps = append(resultOfSteps, controller.RecvData())

// 	controller = NewSyncBridge[int](1, "test-controller-2", 1)
// 	go func() {
// 		defer controller.Done()
// 		controller.SendData(iterator) // 2
// 		iterator++
// 	}()
// 	resultOfSteps = append(resultOfSteps, controller.RecvData())

// 	controller = RunAsync[int]("test-controller-3", func(c *SyncBridge[int]) {
// 		c.SendData(iterator) // 3
// 		iterator++
// 	})
// 	controller.BindCompletionHandler(func(args ...any) {
// 		resultOfSteps = append(resultOfSteps, controller.RecvData())
// 	})
// 	controller.Wait()

// 	controller = RunAsync[int]("test-controller-4", func(c *SyncBridge[int]) {
// 		c.Failed(fmt.Errorf("Error from sender goroutine.")) // passing 4
// 		iterator++
// 	})
// 	controller.BindFailureHandler(func(args ...any) {
// 		assert.NotEmpty(t, controller.Reason())
// 	})
// 	controller.Wait()

// 	controller = RunAsync[int]("test-controller-5", func(c *SyncBridge[int]) {
// 		c.SendData(iterator) // 5
// 		iterator++
// 	})
// 	controller.BindCompletionHandler(func(args ...any) {
// 		resultOfSteps = append(resultOfSteps, controller.RecvData())
// 	})
// 	controller.Wait()

// 	controller = RunAsync[int]("test-controller-6", func(c *SyncBridge[int]) {
// 		c.SendData(iterator) // 6
// 		iterator++
// 		c.SendData(iterator) // 7
// 		iterator++
// 	})
// 	controller.BindCompletionHandler(func(args ...any) {
// 		resultOfSteps = append(resultOfSteps, controller.RecvData())
// 	})
// 	resultOfSteps = append(resultOfSteps, controller.RecvData())
// 	controller.Wait()

// 	assert.Equal(t, len(expectedResultsOfSteps), len(resultOfSteps))
// 	assert.Equal(t, expectedResultsOfSteps[:], resultOfSteps)
// }

// func TestSafeWrapperSlice(t *testing.T) {
// 	defer utils.CleanupLogFiles()

// 	// 10 iterations of checking correctness of NewSafeWrapperSlice.
// 	for i := 0; i < 10; i++ {
// 		safeWrapper := NewSafeWrapperSlice[int]()
// 		iterations := 30
// 		consumer_1_Result := []int{}
// 		awaiter := sync.WaitGroup{}
// 		awaiter.Add(1)
// 		go func() {
// 			defer awaiter.Done()
// 			for i := 0; i < iterations; i++ {
// 				consumer_1_Result = append(consumer_1_Result, safeWrapper.WaitAndGetElement(i))
// 			}
// 		}()
// 		awaiter.Add(1)
// 		consumer_2_Result := []int{}
// 		go func() {
// 			defer awaiter.Done()
// 			for i := 0; i <= iterations-5; {
// 				consumer_2_Result = append(consumer_2_Result, safeWrapper.WaitAndGetRangeOfElements(i, i+5)...)
// 				i += 5
// 			}
// 		}()

// 		for i := 0; i < 15; i++ {
// 			time.Sleep(time.Duration(100 * time.Millisecond))
// 			safeWrapper.AppendNewElement(i)
// 		}
// 		safeWrapper.AppendMany(&[]int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25})
// 		safeWrapper.AppendMany(&[]int{26, 27, 28, 29})

// 		awaiter.Wait()
// 		assert.Equal(t, 30, len(consumer_1_Result))
// 		assert.Equal(t, 30, len(consumer_2_Result))
// 		assert.Equal(t, safeWrapper.WaitAndGetRangeOfElements(0, 30), consumer_1_Result)
// 		assert.Equal(t, safeWrapper.WaitAndGetRangeOfElements(0, 30), consumer_2_Result)
// 	}
// }

// func TestSafeWrapperMap(t *testing.T) {
// 	defer utils.CleanupLogFiles()

// 	safeMap := NewSafeWrapperMap[int, int]()

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			safeMap.AddNewElement(i, i)
// 		}
// 	}()

// 	go func() {
// 		m := make(map[int]int)
// 		for i := 10; i < 20; i++ {
// 			m[i] = i
// 		}
// 		safeMap.AddNewElements(m)
// 	}()

// 	for i := 0; i < 20; i++ {
// 		currentElement, ok := safeMap.WaitAndGetElement(i)
// 		assert.True(t, ok)
// 		assert.Equal(t, i, currentElement)
// 	}

// 	for i := 0; i < 20; i++ {
// 		currentElement, ok := safeMap.TryGetElement(i)
// 		assert.True(t, ok)
// 		assert.Equal(t, i, currentElement)
// 	}
// }
