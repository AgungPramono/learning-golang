package main

import "fmt"

func main() {
	runApp(true)
	fmt.Println("aplikasi tetap berjalan")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}

	//message := recover() //cara menggunakan recover yang salah
	//fmt.Println("Terjadi panic", message)
}

func endApp() {
	fmt.Println("Aplikasi di hentikan secara paksa")
	message := recover() // cara menggunakan recover yang benar
	fmt.Println(message)
}
