package main

import "fmt"

func main() {

	//break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke ", i)
	}

	//continue
	fmt.Println("\ncontinue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke ", i)
	}

}
