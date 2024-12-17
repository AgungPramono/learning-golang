package main

import (
	"fmt"
)

// struct adalah sebuah template data
// artinya struct tidak bisa langsung digunakan
// tapi kita bisa membuat data/object dari struct tersebut
type Customer struct {
	Name, Address string
	phone         int
}

// struct method
// adalah function yg menempel pada sebuah struct
func (customer Customer) info(name string) {
	fmt.Println("hallo", name, " nama kamu adalah", customer.Name)
}

func main() {
	//membuat object dari struct Customer dan mengubah nilai fieldnya
	var agung Customer
	agung.Name = "agung"
	agung.Address = "Indonesia"
	agung.phone = 429572952

	fmt.Println(agung)

	//membuat struct literals
	ahmad := Customer{
		Name:    "ahmad",
		Address: "Indonesia",
		phone:   1234,
	}

	fmt.Println(ahmad)

	joko := Customer{"joko", "Indonesia", 1234}
	fmt.Println(joko)

	//cara memanggil struct method
	agung.info("agung")
	ahmad.info("ahmad")
	joko.info("joko")
}
