package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Terjadi pembagian dengan 0") //membuat error
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := pembagian(100, 0)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hasil :", result)
	}

}
