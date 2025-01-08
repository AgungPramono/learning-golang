package golang_generic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Age int

type TypeNumber interface {
	~int | float32 | float64 | int32 | int64 | int16 | string
}

func GetMin[T TypeNumber](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}
func TestMinNumber(t *testing.T) {
	assert.Equal(t, int(100), GetMin[int](100, 200))
	assert.Equal(t, int64(100), GetMin[int64](100, 200))
	assert.Equal(t, float64(100.1), GetMin[float64](100.1, 200.2))
	assert.Equal(t, Age(23), GetMin[Age](Age(23), Age(25)))
}

func TestTypeInference(t *testing.T) {
	assert.Equal(t, int(100), GetMin(100, 200))
	assert.Equal(t, int64(100), GetMin(int64(100), int64(200)))
	assert.Equal(t, float64(100.1), GetMin(float64(100.1), float64(200.2)))
	assert.Equal(t, Age(23), GetMin(Age(23), Age(25)))
}
