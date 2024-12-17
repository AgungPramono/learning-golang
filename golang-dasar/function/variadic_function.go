package main

import "fmt"

//function bisa menerima banyak inputan/tidak terbatas
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)
	fmt.Println(sumAll(10, 10, 40, 60, 30, 20, 300))

	//jika variabel yg akan dikirim berbentuk slice
	numbers := []int{10, 10, 10, 10, 190}

	//setelah nama variabel di akhiri dengan titik 3 kali
	fmt.Println(sumAll(numbers...))
}
