package main

import "fmt"

//type assertion : kemampuan merubah tipe data menjadi tipe data yg diinginkan
//sering di gunakan ketika bertemu dengan interface kosong

func random() any {
	return 100
}

func main() {
	var result any = random()
	fmt.Println(result)
	//konversi ke tipe string
	//var resultString = result.(string)
	//fmt.Println(resultString)

	//hati2 ketika konversi tipe
	//tipe data tujuan harus sama dengan tipe data tujuan
	//var resultInt = result.(int) // panic
	//fmt.Println(resultInt)

	// agr lebih aman gunakan witch assertion
	switch value := result.(type) {
	case string:
		fmt.Println("tipe data value adalah string", value)
	case int:
		fmt.Println("tipe data value adalah int", value)
	case float64:
		fmt.Println("tipe data value adalah float64", value)
	default:
		fmt.Println("tipe data value tidak diketahui", value)
	}
}
