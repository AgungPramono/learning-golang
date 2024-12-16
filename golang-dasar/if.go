package main

import "fmt"

func main() {
	name := "Andik"

	if name == "Agung" {
		fmt.Println("Hallo Agung")
	} else if name == "Andi" {
		fmt.Println("Hallo Andi")
	} else {
		fmt.Println("Hallo", "siapa kamu?")
	}

	//if short statement
	var newName string = "ari"
	if length := len(newName); length > 5 {
		fmt.Println("Nama kamu terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
