package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}

func TestLogJson(t *testing.T) {
	logJson("agung")
	logJson(1)
	logJson(true)
	logJson([]string{"Golang", "Java", "C++"})
}
