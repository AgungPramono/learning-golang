package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	//membuat context
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	//ambil context value
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextD.Value("b"))
	//fmt.Println(contextF.Value("b"))
}

func createCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done(): // cek ager tidak terjadi goroutine leak
				return
			default:
				destination <- counter
				counter++
			}

		}
	}()

	return destination
}

func TestCounter(t *testing.T) {

	fmt.Println("goroutine number ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent) //buat child context dg konfigurasi cancel

	destination := createCounter(ctx)

	fmt.Println("goroutine number ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter : ", n)
		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)
	fmt.Println("goroutine number ", runtime.NumGoroutine())
}
