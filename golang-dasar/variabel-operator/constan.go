package main

import "fmt"

func main() {

	// membuat constan
	// tidak bisa diubah nilainya
	const firstName = "budiono"
	const lastName = "siregar"

	//multiple constan
	const (
		namaDepan    string = "budiono"
		namaBelakang        = "siregar"
	)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

}
