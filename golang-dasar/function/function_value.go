package main

import "fmt"

/**
function yang bisa digunakan sebagai value
 atau disimpan dalam sebuah variabel
*/

func getName(name string) string {
	return "your name is " + name
}

func main() {
	contoh := getName
	misal :=  getName

	fmt.Println(contoh("Agung"))
	fmt.Println(misal("Budiono"))
}
