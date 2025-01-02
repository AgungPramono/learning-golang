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

func TestArrayComplexTest(t *testing.T) {
	customer := Customer{
		FirstName:  "joko",
		MiddleName: "pitono",
		LastName:   "negoro",
		Age:        18,
		Married:    false,
		Hobbies:    []string{"Reading", "Gaming", "Coding"},
		Addresses: []Address{
			Address{
				"jalan kenangan", "Indonesia", "64451",
			},
			Address{
				"jalan kebeneran", "Indonesia", "68832",
			},
			Address{
				"jalan kebeneran", "Indonesia", "68832",
			},
		},
	}

	jsonCustomer, _ := json.MarshalIndent(customer, "", " ")
	fmt.Println(string(jsonCustomer))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{
						 "FirstName": "joko",
						 "MiddleName": "pitono",
						 "LastName": "negoro",
						 "Age": 18,
						 "Married": false,
						 "Hobbies": [
						  "Reading",
						  "Gaming",
						  "Coding"
						 ],
						 "Addresses": [
						  {
						   "Street": "jalan kenangan",
						   "Country": "Indonesia",
						   "PostalCode": "64451"
						  },
						  {
						   "Street": "jalan kebaikan",
						   "Country": "Indonesia",
						   "PostalCode": "68832"
						  },
						  {
						   "Street": "jalan kebenaran",
						   "Country": "Indonesia",
						   "PostalCode": "68832"
						  }
						 ]
					}`
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
	fmt.Println(customer.Addresses)
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{
						   "Street": "jalan kenangan",
						   "Country": "Indonesia",
						   "PostalCode": "64451"
						  },
						  {
						   "Street": "jalan kebaikan",
						   "Country": "Indonesia",
						   "PostalCode": "68832"
						  },
						  {
						   "Street": "jalan kebenaran",
						   "Country": "Indonesia",
						   "PostalCode": "68832"
						  }
						 ]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

func TestOnlyJsonArrayComplexTest(t *testing.T) {
	addresses := []Address{
		{
			"jalan kenangan", "Indonesia", "64451",
		},
		Address{
			"jalan kebeneran", "Indonesia", "68832",
		},
		Address{
			"jalan kebeneran", "Indonesia", "68832",
		},
	}

	jsonAddreses, _ := json.MarshalIndent(addresses, "", " ")
	fmt.Println(string(jsonAddreses))
}
