package go_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPools(t *testing.T) {

	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Abim")
	pool.Put("Dhanu")
	pool.Put("Ejas")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
			// time.Sleep(1 * time.Second)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Finish")
}
