package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// function untuk unit test harus di awali dengan "Test"
// untuk menjalankan test gunakan perintah : go test
// go test -v : melihat function apa saja yg dijalankan test
// go test -v -run <NamaFunction> : menjalankan test untuk function tertentu
// go test -v ./... : menjalanlan test dari root module

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWorld(t *testing.T) {
	result := HellowWorld("agung")

	if result != "hello agung" {
		//error test
		//t.Fail() // ketika dijalankan tetap akan di lanjutkan ke bawah
		t.Error("result must be \"hello agung\"")
	}
	fmt.Println("TestHelloWorld done")
}

func TestHellodAhmad(t *testing.T) {
	result := HellowWorld("ahmad")

	if result != "hello ahmad" {
		//error test
		//t.FailNow() // ketika dijalankan tidak akan di lanjutkan ke bawah
		t.Fatalf("result must be \"hello ahmad\"")
	}
	fmt.Println("TestHelloAhmad done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HellowWorld("agung")
	assert.Equal(t, "hello agung", result, "result must be \"hello agung\"")
	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldRequired(t *testing.T) {
	result := HellowWorld("agung")
	require.Equal(t, "hello agung", result, "result must be \"hello agung\"")
	fmt.Println("TestHelloWorld with required done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skip this test on windows")
	}
	result := HellowWorld("agung")
	require.Equal(t, "hello agung", result, "result must be \"hello agung\"")
}
