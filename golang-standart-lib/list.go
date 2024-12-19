package main

import (
	"container/list"
	"fmt"
)

func main() {

	//deklarasi double linkedlist
	var data *list.List = list.New()

	//menambahkan data kedalama linkedlist
	data.PushBack("Java")
	data.PushBack("Golang")
	data.PushBack("Php")

	var head = data.Front() // mengambil data paling depan
	fmt.Println(head.Value)

	next := head.Next() //ambil data selanjutnya
	fmt.Println(next.Value)

	next = next.Next() //ambil data selanjutnya
	fmt.Println(next.Value)

	//ambil data menggunakan perulangan
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
