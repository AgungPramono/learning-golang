package main

import "fmt"

func main() {
	name := "Andik"

	switch name {
	case "Agung":
		fmt.Println("Halo, Agung!")
	case "Ahmad":
		fmt.Println("Halo, Ahmad!")
	default:
		fmt.Println("Halo, ", name)
	}

	//switch dengan short statemen
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
