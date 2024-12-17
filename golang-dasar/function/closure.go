package main

import "fmt"

func main() {

	var getYourIdentity = func(name string) (string, int) {
		var group string
		var id int
		if len(name) > 3 {
			group = "group 1"
			id = 11
		} else {
			group = "group 2"
			id = 11
		}
		return group, id
	}

	newGroup, newId := getYourIdentity("Gopher")
	fmt.Println(newGroup)
	fmt.Println(newId)

	counter := 0

	increment := func() {
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)

}
