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
