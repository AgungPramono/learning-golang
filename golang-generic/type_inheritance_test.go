package golang_generic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (manager *MyManager) GetName() string {
	return manager.Name
}

func (manager *MyManager) GetManagerName() string {
	return manager.Name
}

type VicePresiden interface {
	GetName() string
	GetVicePresidenName() string
}

type MyVicePresiden struct {
	Name string
}

func (myVicePresiden *MyVicePresiden) GetName() string {
	return myVicePresiden.Name
}

func (myVicePresiden *MyVicePresiden) GetVicePresidenName() string {
	return myVicePresiden.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "agung", GetName[Manager](&MyManager{Name: "agung"}))
	assert.Equal(t, "budi", GetName[VicePresiden](&MyVicePresiden{Name: "budi"}))
}
