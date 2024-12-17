package main

import "fmt"

func main() {

	/* counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}

	fmt.Println("selesai")
	*/

	//init & pos statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke ", counter)
	}

	//for Range
	//dilakukan untuk mengakses data collection

	names := []string{"Ardi", "Nita", "Agung"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//for range
	for index, name := range names {
		fmt.Println("index ", index, "=", name)
	}

	//jika tidak ingin ada variabel index maka ganti dengan _
	for _, name := range names {
		fmt.Println(name)
	}

}
