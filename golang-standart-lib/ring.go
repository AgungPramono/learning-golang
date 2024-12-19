package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	data := ring.New(10)

	//data.Value = "value 1" //isi data secara manual
	//
	//data = data.Next()
	//data.Value = "value 2"

	//isi data menggunakan perulangan
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value" + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})

}
