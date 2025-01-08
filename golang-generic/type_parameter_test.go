package golang_generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length[string]("agung")
	assert.Equal(t, "agung", result)

	var intResult int = Length[int](12)
	assert.Equal(t, 12, intResult)

}
