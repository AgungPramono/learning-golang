package main

import "fmt"

func main() {
	name := getHello("Agung")
	fmt.Println(name)

	fmt.Println(getHello("Budiono"))
	fmt.Println(getHello("Ahmad"))
	fmt.Println(getHello("Dedi"))
}

func getHello(name string) string {
	hello := "Hello" + name
	return hello
}
