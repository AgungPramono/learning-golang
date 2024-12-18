package main

import "fmt"

//nil akan mengembalikan nilai kosong
//nil hanya bisa digunakan pada type data
//interface, map , function, slice dan channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	nameMap := NewMap(" ")
	if nameMap == nil {
		fmt.Println("data map masih kosong")
	} else {
		fmt.Println(nameMap["name"])
	}
}
