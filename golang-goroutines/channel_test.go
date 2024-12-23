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

// function hanya bisa untuk mengirim data melalui channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	//data := <-channel //tidak bisa
	channel <- "Pemrograman Golang Untuk Pemula"
}

// function hanya bisa untuk mengambil data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	//channel <- "Pemrograman java" //tidak bisa
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
	defer close(channel)
}

func TestChannelWithBuffer(t *testing.T) {
	channel := make(chan string, 3) //buat channel dg buffer 3
	defer close(channel)

	go func() {
		channel <- "Belajar Java"
		channel <- "Belajar Golang"
		channel <- "Belajar Js"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(2 * time.Second)

	fmt.Println("Selesai mengirim data ke channel")

}
