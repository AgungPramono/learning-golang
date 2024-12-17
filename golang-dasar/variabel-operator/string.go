package main

import "fmt"

func main() {

	fmt.Println("selamat datang")
	fmt.Println("ini adalah program pertama saya")

	//menghitung panjang string
	fmt.Println("panjang string adalah", len("ini adalah string"))
	fmt.Println("panjang string adalah", len("panjang string"))

	//mengambil karakter (mengembalikan nilai byte)
	fmt.Println("ini adalah string"[0])
	fmt.Println("ini adalah string"[1])
}
