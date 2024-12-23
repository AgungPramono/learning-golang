package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() //menjalankan proses menggunakan goroutine
	fmt.Println("Ow")

	time.Sleep(1 * time.Second) //1 detik
}

func DisplayNumber(number int) {
	fmt.Println("Display :", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
