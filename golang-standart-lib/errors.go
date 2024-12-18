package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found")
)

func getById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "agung" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := getById("eko")

	//cara pengecekkan jenis error
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("NotFoundError")
		} else {
			fmt.Println("Unknown Error")
		}
	} else {
		fmt.Println("Success")
	}
}
