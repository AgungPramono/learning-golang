package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//cara manual pointer
	//var alamat1 *Address = &Address{}
	//menggunakan new
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
