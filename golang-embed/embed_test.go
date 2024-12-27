package golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//embed file txt kedalam variable string
//hanya bisa di lakukan di luar function

//go:embed version.txt
var version string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}

// embed file binary(gambar)
//
//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("log_new.png", logo, fs.ModePerm)
	if err != nil {
		return
	}
}
