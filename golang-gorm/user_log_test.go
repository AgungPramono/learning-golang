package golang_gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 15; i++ {
		userLog := UserLog{
			UserId: strconv.Itoa(i),
			Action: "Test Action",
		}
		result := Db().Create(&userLog)
		assert.Nil(t, result.Error)
		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}
}

func TestSaveOrUpdate(t *testing.T) {
	userLog := UserLog{
		UserId: "15",
		Action: "Test Action 15",
	}
	result := Db().Save(&userLog)
	assert.Nil(t, result.Error)
	assert.Equal(t, "Test Action 15", userLog.Action)

	userLog.UserId = "2"
	userLog.Action = "Test Action 15 update"
	err := Db().Save(&userLog).Error
	assert.Nil(t, err)
}
