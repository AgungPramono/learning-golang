package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var address1 Address = Address{"Nganjuk", "Jatim", "Indonesia"}
	var address2 *Address = &address1 //pointer

	address2.City = "Surabaya"
	fmt.Println(address1)
	fmt.Println(address2)

	//address2 = &Address{"Sidoarjo", "Jatim", "Indonesia"}

	//jika menggunakan operator *:
	// semua yg ditunjuk address2
	// akan ikut berubah
	*address2 = Address{"Malang", "Jatim", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
