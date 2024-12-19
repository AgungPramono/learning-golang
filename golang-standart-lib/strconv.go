package main

import (
	"fmt"
	"strconv"
)

func main() {

	result, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	//format binary
	binary := strconv.FormatInt(192, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(2345)
	fmt.Println(stringInt)
}
