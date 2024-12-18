package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

func (e *notFoundError) Error() string {
	return e.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "agung" {
		return &notFoundError{"data not found"}
	}
	return nil
}

func main() {
	err := saveData("agung", nil)

	if err != nil {
		//if validationErr, ok := err.(*validationError); ok { // konversi menggunakan type assertion
		//	fmt.Println("validation error:", validationErr.Error())
		//} else if notFoundErr, ok := err.(*notFoundError); ok {
		//	fmt.Println("not found error:", notFoundErr.Error())
		//} else {
		//	fmt.Println("unknown error:", err.Error())
		//}

		//menggunakan switch case
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error", finalError.Error())
		default:
			fmt.Println("unknown error", finalError.Error())
		}
	} else {
		fmt.Println("data saved successfully")
	}

}
