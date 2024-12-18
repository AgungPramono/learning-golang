package main

import "fmt"

type Man struct {
	Name string
}

// data tidak akan berubah karena yang parameter man berupa value
// yang sudah di duplikat
//func (man Man) Married() {
//	man.Name = "Mr." + man.Name
//}

// agar data tidak di duplikat
// atau data aslinya yg diubah
// gunakan pointer method
func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	joko := Man{"joko"}
	joko.Married()

	fmt.Println(joko)

}
