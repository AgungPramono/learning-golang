package main

import (
	"fmt"
)

func main() {

	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// konversi tipe data string
	var nama = "Rizky"
	var e uint8 = nama[0]
	var estring = string(e)

	fmt.Println(estring)
}
