package main

import "fmt"

// buat interface kosong
func Test() any {
	//return 1
	return true
	//return "Tests"
}

func main() {
	var kosong any = Test()
	fmt.Println(kosong)

}
