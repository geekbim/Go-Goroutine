package go_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU ", totalCpu)

	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread ", totalThreads)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine ", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU ", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread ", totalThreads)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine ", totalGoroutine)

	group.Wait()
}
