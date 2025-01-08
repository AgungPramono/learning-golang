package golang_generic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Number interface {
	int | float32 | float64 | int32 | int64 | int16 | string
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}
func TestMin(t *testing.T) {
	assert.Equal(t, int(100), Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](100, 200))
	assert.Equal(t, float64(100.1), Min[float64](100.1, 200.2))
}
