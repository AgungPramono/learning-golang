package main

import "fmt"

func main() {

	var numberA int = 3
	//referencing variabel ke variabel pointer
	var numberB *int = &numberA

	fmt.Println(numberA)
	fmt.Println(numberB)

	//deferencing pointer ke variabel
	fmt.Println(*numberB)
	fmt.Println(numberB)

	//ketika nilai variabel di ubah maka nilai variabel pointer yg menunjuk
	// ke variabel tsb juga akan ikut berubah

	numberA = 10

	fmt.Println(numberA)
	fmt.Println(&numberA)

	fmt.Println(*numberB)
	fmt.Println(numberB)

	if *numberB == numberA {
		fmt.Println("numberA is equal to numberB")
	}

}
