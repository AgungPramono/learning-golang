package golang_gorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	connection := OpenConnection()
	assert.NotNil(t, connection)
}
