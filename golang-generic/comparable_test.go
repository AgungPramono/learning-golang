package golang_generic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSame[T comparable](value1 T, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	assert.True(t, true, IsSame[string]("ahmad", "agung"))
	assert.True(t, true, IsSame[int](21, 100))
	assert.True(t, true, IsSame[bool](true, false))
}
