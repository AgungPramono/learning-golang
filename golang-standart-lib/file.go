package main

import (
	"bufio"
	"io"
	"os"
)

// buat file dan tulis ke dalam file
func createNewFile(fileName string, message string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	//tutup file
	defer file.Close()
	file.WriteString(message)
	return nil
}

// tambah baris baru ke file
func addToFile(fileName string, message string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	//tutup file
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	//proses baca file
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	//createNewFile("sample.log", "ini sample log")
	//result, _ := readFile("sample.log")
	//fmt.Println(result)

	addToFile("sample.log", "add new line to file")
}
