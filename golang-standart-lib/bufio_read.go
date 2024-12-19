package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	data := "Pemrograman Golang\nuntuk pemula\nlengkap\n"

	input := strings.NewReader(data)
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

}
