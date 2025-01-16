package golang_gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-gorm/model"
	"gorm.io/gorm/clause"
	"strconv"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 15; i++ {
		userLog := model.UserLog{
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
	userLog := model.UserLog{
		UserId: "15",
		Action: "Test Action 15",
	}
	result := Db().Save(&userLog) //insert
	assert.Nil(t, result.Error)
	assert.Equal(t, "Test Action 15", userLog.Action)

	userLog.UserId = "2"
	userLog.Action = "Test Action 15 update"
	err := Db().Save(&userLog).Error //update
	assert.Nil(t, err)
}

func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
	user := model.User{
		Id: "99",
		Name: model.Name{
			FirstName: "Joko",
		},
	}
	err := Db().Save(&user).Error //insert
	assert.Nil(t, err)

	user.Name.FirstName = "joko Pitono 99"
	err = Db().Save(&user).Error //update
	assert.Nil(t, err)
}

func TestConflict(t *testing.T) {
	user := model.User{
		Id: "100",
		Name: model.Name{
			FirstName: "Joko 100",
		},
	}
	err := Db().Clauses(clause.OnConflict{
		UpdateAll: true},
	).Create(&user).Error //insert
	assert.Nil(t, err)
}
