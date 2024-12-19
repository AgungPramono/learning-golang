package main

import (
	"fmt"
	"regexp"
)

func main() {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	var regex = regexp.MustCompile(emailRegex)
	fmt.Println(regex.MatchString("agung@mail.com"))
	fmt.Println(regex.MatchString("agungmail@.com"))
	fmt.Println(regex.MatchString("agung@ub.ac.id"))
}
