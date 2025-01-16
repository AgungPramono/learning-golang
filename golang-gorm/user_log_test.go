package golang_gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 15; i++ {
		userLog := UserLog{
			UserId: "1",
			Action: "Test Action",
		}
		result := Db().Create(&userLog)
		assert.Nil(t, result.Error)
		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}
}
