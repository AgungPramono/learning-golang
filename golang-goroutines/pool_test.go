package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{}

	pool.Put("Golang")
	pool.Put("Java")
	pool.Put("C++")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data) // kembalikan lagi data ke pool
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
