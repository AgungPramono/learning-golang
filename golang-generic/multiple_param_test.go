package golang_generic

import (
	"fmt"
	"testing"
)

func MultipleParamTest[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultiParameter(t *testing.T) {
	MultipleParamTest[string, int]("agung", 12)
	MultipleParamTest[int, string](100, "agung")
}
