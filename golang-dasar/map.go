package main

import "fmt"

func main() {

	//mendeklarasikan map
	/* var person map[string]string = map[string]string{}

	//isi map
	person["kode"] = "001"
	person["name"] = "agung"

	person["kode"] = "002"
	person["name"] = "ahmad"
	*/
	//deklarasi map menggunakan blok
	person := map[string]string{
		"kode": "001",
		"name": "agung",
	}

	fmt.Println(person["kode"])
	fmt.Println(person["name"])
	fmt.Println(person)
	println("panjang map", len(person))

	//fungsi yang dalam map
	book := make(map[string]string)
	book["title"] = "Pemrograman Java"
	book["author"] = "Agung"
	book["err"] = "salah"

	fmt.Println(book)

	//hapus
	delete(book, "err")

	fmt.Println(book)

}
