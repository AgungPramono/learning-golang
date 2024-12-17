package main

import "fmt"

func main() {

	//membuat array
	names := []string{"John", "Doni", "Jane", "David", "Alice", "Ahmad"}

	//array[low:high]

	//membuat slice 1
	slice1 := names[4:6]
	fmt.Println(slice1)

	//buat array dari index 0 sampai sebelum index high
	slice2 := names[:2]
	fmt.Println(slice2)

	//buat slice dimulai dari index low sampai akhir index array
	slice3 := names[3:]
	fmt.Println(slice3)

	//buat slice dari awal array dimulai index 0 sampai index akhir array
	slice4 := names[:]
	fmt.Println(slice4)

	slice5 := names[2:4]
	fmt.Println(slice5)

	slice6 := names[3:4]
	fmt.Println(slice6)

	slice7 := names[1:3]
	fmt.Println(slice7)

	//function di slice

	//panjang slice
	fmt.Println(len(slice1))
	//kapasitas slice
	fmt.Println(cap(slice7))

	days := []string{"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}

	daysSlice := days[0:2]
	// fmt.Println(daysSlice)

	//mengganti nilai di slice
	daysSlice[0] = "senin baru"
	daysSlice[1] = "selasa baru"
	fmt.Println(daysSlice)
	fmt.Println(days) //[senin baru selasa baru rabu kamis jum'at sabtu minggu]

	//append
	//append() menambahkan nilai di akhir array
	daysSlice2 := append(daysSlice, "libur baru")
	daysSlice2[0] = "selasa lama"
	// fmt.Println(daysSlice)
	// fmt.Println(daysSlice2)
	fmt.Println(days)

	//buat slice dari awal
	//buat slice dengan tipe string dengan panjang 2 dan kapasitas 5
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Agung"
	newSlice[1] = "Agung"
	// newSlice[2] = "Agung" // error, harus menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//append slice baru
	newSlice2 := append(newSlice, "Agung")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// cek apakah newSlice2 masih menggunakan array yg sama
	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	//copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//copy lagi ke slice baru
	newToSlice := make([]string, len(toSlice), cap(toSlice))
	copy(newToSlice, toSlice)

	fmt.Println(newToSlice)

	//beda array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5} //array
	iniSlice := []int{1, 2, 3, 4, 5}    //slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
