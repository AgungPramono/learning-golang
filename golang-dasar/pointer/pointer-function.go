package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	//address := Address{}
	//ChangeCountryToIndonesia(address)
	//
	//fmt.Println(address) //nilai masih kosong

	//agar addres tidak kosong gunakan pointer *
	var address3 *Address = &Address{}
	ChangeCountryToIndonesia(address3)

	fmt.Println(address3)
}
