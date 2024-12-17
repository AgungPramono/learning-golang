package main

import (
	"fmt"
)

// buatkan type atau alis function agar jika jumlah parameter banyak
// keyword type
// maka tidak perlu menulis banyak parameter
type Filter func(string) string

func messageWithFilter(message string, filter func(string) string) {
	filteredMessage := filter(message)
	fmt.Println(filteredMessage)
}

func simpleMessageFilter(message string, filter Filter) {
	filteredMessage := filter(message)
	fmt.Println(filteredMessage)
}

func wordFilter(message string) string {
	if message == "babi" || message == "anjing" {
		return "***"
	} else {
		return message
	}
}

func main() {

	messageWithFilter("hallo", wordFilter)
	simpleMessageFilter("babi", wordFilter)

	message := wordFilter
	messageWithFilter("babi", message)
	simpleMessageFilter("anjing", message)

}
