package main

import "fmt"

func main() {
	name, age := getIdentity("agung", 18)
	fmt.Println(name, age)

	//ignore return value
	//ganti dengan _ untuk return yang tidak butuh di ambil nilai returnnya
	//_ = getIdentity("agung", 18)
	name, _ = getIdentity("budiono", 25)
	fmt.Println(name)
}

func getIdentity(name string, age int) (string, int) {
	return name, age
}
