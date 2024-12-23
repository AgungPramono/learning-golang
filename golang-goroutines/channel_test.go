package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) //buat channel

	go func() {
		time.Sleep(2 * time.Second)
		//input data ke channel
		channel <- "Pemrograman Golang"
		fmt.Println("selesai mengirim data ke channel")
	}()

	//ambil data dari channel
	data := <-channel
	fmt.Println(data)
	defer close(channel)
	time.Sleep(5 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	defer close(channel)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Pemrograman Golang Untuk Pemula"
}
