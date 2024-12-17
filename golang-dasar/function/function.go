package main

import "fmt"

func main() {
	sayHello()
	sayHelloTo("Agung", "Budiono")
}

func sayHello() {
	fmt.Println("Hello")
}

// function parameter
func sayHelloTo(name string, lastName string) {
	fmt.Println("Hello", name, lastName)
}
