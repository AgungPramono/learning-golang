// tipe data yang berisikan kumpulan data dengan tipe yang sama
// buat array perlu menentukan ukuran yang bisa di tampung oleh array tersebut
// index array di mulai dari 0

package main

import "fmt"

func main() {

	// buat array
	var name [3]string

	//mengisi array
	name[0] = "agung"
	name[1] = "budi"
	name[2] = "candra"
	// name[4] = "adi" // error karena index array di atas 3

	// tampilkan isi masing2 array
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])
	// fmt.Println(name[4])

	//mengubah isi array index 2
	name[2] = "eko"
	fmt.Println(name[2])

	//membuat dan mengisi array secara langsung
	var daftarNama = [3]string{
		"agung",
		"budi",
		"candra", //jika enter wajib tambah koma di akhir
	}

	// jumlah array tidak ditentukan langsung
	// var daftarNama = [...]string{}

	fmt.Println(daftarNama)

	// function array

	//menghitung panjang array
	var panjangArray int = len(daftarNama)
	fmt.Println(panjangArray)

}
