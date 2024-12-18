package main

import "fmt"

// cara membuat interface
type HasName interface {
	GetName() string
}

// sebelum implementasi interface buat dulu struct
type Person struct {
	Name string
}

// kemudian buat implementasi interfacenya
func (person Person) GetName() string {
	return person.Name
}

func SayHello(name HasName) {
	fmt.Println("Your name is :", name.GetName())
}

func main() {

	// panggil impl interface
	person := Person{Name: "Agung"}
	SayHello(person)

}
