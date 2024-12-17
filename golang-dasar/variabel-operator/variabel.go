package main

import (
	"fmt"
)

func main() {

	// membuat variabel
	//jika tidak langsung di inisialisasi tipe data harus disebutkan
	var name string

	name = "budiono siregar"
	fmt.Println(name)

	//ubah nilai variabel
	name = "kapal laut"
	fmt.Println(name)

	//membuat variabel dengan tipe data
	var umur int = 25
	fmt.Println(umur)

	//tidak perlu menggunakan var
	//asal menggunakan :=
	//hanya bisa digunakan di awal
	nama := "budiono"
	fmt.Println(nama)

	//deklarasi multiple variabel
	var (
		namaLengkap = "Budiono"
		cita2       = "Kapal laut"
	)

	fmt.Println(namaLengkap)
	fmt.Println(cita2)

}
