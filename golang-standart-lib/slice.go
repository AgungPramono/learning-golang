package main

import (
	"fmt"
	"slices"
)

func main() {

	names := []string{"John", "Paul", "George"}
	values := []int{21, 34, 56}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Paul"))
	fmt.Println(slices.Index(names, "John"))
	fmt.Println(slices.Index(names, "George"))

}
