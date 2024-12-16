package main

import "fmt"

func main() {

	// membuat tipe data
	type NoKtp string

	var ktpBudi NoKtp = "111111111"

	var contoh NoKtp = "1234567890"
	var contohKtp NoKtp = NoKtp(ktpBudi)

	fmt.Println(contoh)
	fmt.Println(contohKtp)

}
