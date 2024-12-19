package main

import (
	"fmt"
	"sort"
)

type User struct {
	Nama string
	Age  int
}

type UserSlice []User

// buat kontrak untuk proses sorting
func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	users := []User{
		{"Agung", 18},
		{"Budi", 17},
		{"Rizky", 12},
		{"Yono", 34},
		{"Dedi", 56},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

}
