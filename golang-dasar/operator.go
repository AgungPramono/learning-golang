package main

import "fmt"

func main() {

	var a = 10
	var b = 10
	var c = a + b
	var d = c + c

	fmt.Println(d)

	// augmented assignments
	c += c
	fmt.Println(c)

	//unary operator
	var e = 10
	e++ // e=e+1
	fmt.Print(e)

	e-- // e = e-1;
	fmt.Println(e)

	//operasi perbandingan
	var nama1 = "Rahmad"
	var nama2 = "Rahmad"

	var result bool = nama1 == nama2

	fmt.Println(result)

	//operasi Boolean

}
