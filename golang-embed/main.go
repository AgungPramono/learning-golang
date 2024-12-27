package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed test/version.txt
var version string

//go:embed test/logo.png
var logo []byte

//go:embed test/files/file-1.txt
//go:embed test/files/file-2.txt
//go:embed test/files/file-3.txt
var files embed.FS

//go:embed test/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("log_new2.png", logo, fs.ModePerm)
	if err != nil {
		return
	}

	dir, _ := path.ReadDir("test/files")
	for _, entry := range dir {

		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("test/files/" + entry.Name())
			fmt.Println("conten: ", string(content))
		}
	}
}
