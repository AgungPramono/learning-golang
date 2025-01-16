package golang_gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-gorm/model"
	"testing"
)

func TestCreateUserAddresses(t *testing.T) {
	user := model.User{
		Id:       "20",
		Password: "rahasia 20",
		Name: model.Name{
			FirstName: "Jumadi 20",
		},
		Wallet: model.Wallet{
			Id:      "3",
			UserId:  "22",
			Balance: 1_000_000,
		},
		Addresses: []model.Address{
			{
				UserId:  "20",
				Address: "Madiun",
			},
			{
				UserId:  "20",
				Address: "Nganjuk",
			},
		},
	}

	err := Db().Save(&user).Error
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var userPreload []model.User
	err := Db().Model(&model.User{}).Preload("Addresses").Joins("Wallet").Find(&userPreload).Error
	assert.Nil(t, err)
}

func TestTakePreloadJoinOneToMany(t *testing.T) {
	var user model.User
	err := Db().Model(&model.User{}).Preload("Addresses").Joins("Wallet").Take(&user, "users.id=?", "22").Error
	assert.Nil(t, err)
}

func TestBelongsToAddressUser(t *testing.T) {
	fmt.Println("Preload")
	var addresses []model.Address
	err := Db().Model(&model.Address{}).Preload("User").Find(&addresses).Error
	assert.Nil(t, err)
	fmt.Println(len(addresses))

	fmt.Println("Join")
	addresses = []model.Address{}
	err = Db().Model(&model.Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
}
