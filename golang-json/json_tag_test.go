package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P-001",
		Name:     "Produk 001",
		ImageURL: "http://toko.com/image/product_001.png",
	}

	jsonProduct, _ := json.Marshal(product)
	fmt.Println(string(jsonProduct))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"Id":"P-001","name":"Produk 001","image_url":"http://toko.com/image/product_001.png"}`
	jsonBytes := []byte(jsonString)

	product := Product{}
	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
