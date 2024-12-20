package helper

import "testing"

// function untuk unit test harus di awali dengan "Test"
// untuk menjalankan test gunakan perintah : go test
// go test -v : melihat function apa saja yg dijalankan test
// go test -v -run <NamaFunction> : menjalankan test untuk function tertentu
// go test -v ./... : menjalanlan test dari root module

func TestHelloWorld(t *testing.T) {
	result := HellowWorld("agung")
	if result != "hello agung" {
		//error test
		panic("Result is not hello Eko")
	}
}

func TestHellodAhmad(t *testing.T) {
	result := HellowWorld("ahmad")
	if result != "hello ahmad" {
		//error test
		panic("Result is not hello Eko")
	}
}
