package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Nganjuk", "Jatim", "Indonesia"}
	address2 := address1 // value address1 dicopy ke address2

	fmt.Println(address1)
	fmt.Println(address2)

	address2.City = "Kediri" //ganti kota
	fmt.Println(address1)    // value tidak akan berubah
	fmt.Println(address2)    // value berubah

	//POINTER
	address3 := Address{"Solo", "Jateng", "Indonesia"}
	address4 := &address3 // address4 adalah pointer ke address3
	//var address4 *address = &address3

	fmt.Println(address3)
	fmt.Println(address4)

	address4.Country = "Semarang" //jika address4 di ubah
	fmt.Println(address3)         // maka address3 akan ikut berubah
	fmt.Println(address4)

	//OPERATOR *

}
