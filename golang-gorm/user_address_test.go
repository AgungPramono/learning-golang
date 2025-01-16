package golang_gorm

import (
	"github.com/stretchr/testify/assert"
	"golang-gorm/model"
	"testing"
)

func TestCreateUserAddresses(t *testing.T) {
	user := model.User{
		Id:       "22",
		Password: "rahasia 22",
		Name: model.Name{
			FirstName: "Ahmad 22",
		},
		Wallet: model.Wallet{
			Id:      "3",
			UserId:  "22",
			Balance: 1_000_000,
		},
		Address: []model.Address{
			{
				UserId:  "22",
				Address: "Surabaya",
			},
			{
				UserId:  "22",
				Address: "Kediri",
			},
		},
	}

	err := Db().Create(&user).Error
	assert.Nil(t, err)
}
