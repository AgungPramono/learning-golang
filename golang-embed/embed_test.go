package golang_embed

import (
	"embed"
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

// embed multiple file
//
//go:embed files/file-1.txt
//go:embed files/file-2.txt
//go:embed files/file-3.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/file-1.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/file-2.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/file-3.txt")
	fmt.Println(string(c))
}
