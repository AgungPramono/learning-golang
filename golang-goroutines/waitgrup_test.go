package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("hello world")
	time.Sleep(1 * time.Second)
}

func TestWaitGrup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i <= 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Complete")
}
