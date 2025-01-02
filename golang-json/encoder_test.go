package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestJsonStreamEncoder(t *testing.T) {
	writer, _ := os.Create("cutomer_out.json")
	jsonEncoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Joko",
		MiddleName: "waluyo",
		LastName:   "jati",
	}

	jsonEncoder.Encode(customer)
}
