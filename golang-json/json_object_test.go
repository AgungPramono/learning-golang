package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Budiono",
		MiddleName: "Siregar",
		LastName:   "Kapal Laut",
		Age:        24,
		Married:    true,
	}

	jsonCustomer, _ := json.Marshal(customer)

	fmt.Println(string(jsonCustomer))
}
