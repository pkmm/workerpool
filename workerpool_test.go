package workerpool

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func toSleep(i int) {
	time.Sleep(time.Duration(i) * time.Second)
	wg.Done()
}

func TestWorkerPool_Simple(t *testing.T) {
	t.Parallel()
	var tt sync.WaitGroup
	jobCount := 86
	wp := NewWorkerPool(17)

	for i := 0; i < 10; i++ {
		wp.Start()
		wp.Stop()
	}

	wp.Start()

	go func() {
		for {
			select {
			case <-time.Tick(time.Second * 1):
				t.Logf("=====>!!! worker count %d\n", wp.GetActiveWorkerCount())
			}
		}
	}()

	tt.Add(jobCount)
	t.Log("begin\n")
	go func() {
		for i := 0; i < jobCount; i++ {
			y := i
			wp.Execute(func() {
				rd := rand.Intn(8) + 1
				if y > 40 {
					rd = 1
				}
				time.Sleep(time.Duration(rd) * time.Second)
				if rd == 6 {
					tt.Done()
					panic("panic in user func.")
				}
				t.Logf("task_id %d, 休眠的的时间: %d\n", y+1, rd)
				tt.Done()
			})
		}
	}()

	tt.Wait()
	wp.Stop()
	t.Log("over\n")
}
