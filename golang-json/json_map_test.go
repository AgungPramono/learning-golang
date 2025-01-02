package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapTest(t *testing.T) {
	jsonString := `{"id":"P-001","name":"Produk 001","price":230000,"image_url":"http://toko.com/image/product_001.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
	fmt.Println(result["image_url"])
}

func TestJsonMapDecode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P-001",
		"name":      "Produk 001",
		"price":     230000,
		"image_url": "http://toko.com/image/product_001.png",
	}

	jsonBytes, _ := json.Marshal(product)
	fmt.Println(string(jsonBytes))
}
