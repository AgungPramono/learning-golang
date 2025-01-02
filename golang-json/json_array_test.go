package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "joko",
		MiddleName: "pitono",
		LastName:   "negoro",
		Age:        18,
		Married:    false,
		Hobbies:    []string{"Reading", "Gaming", "Coding"},
	}

	jsonCustomer, _ := json.Marshal(customer)
	fmt.Println(string(jsonCustomer))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"John","MiddleName":"Doe","LastName":"Doe","Age":18,"Married":false,"Hobbies":["Reading","Writing","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}
