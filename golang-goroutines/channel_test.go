package golang_goroutines

import (
	"fmt"
	"strconv"
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

// range channel
// menerima data dari channel yg jumlahnya tidak tentu
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke :" + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println("Menenerima data:", data)
	}

	fmt.Println("Selesai mengirim data ke channel")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	//select {
	//case data := <-channel1:
	//	fmt.Println("data dari channel 1:", data)
	//case data := <-channel2:
	//	fmt.Println("data dari channel 2:", data)
	//}

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2:", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}

}
