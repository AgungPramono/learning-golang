package golang_gorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		Id:      "1",
		UserId:  "3",
		Balance: 40_000_000,
	}
	err := Db().Create(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRelationUserWallet(t *testing.T) {
	var user User
	result := Db().Model(&User{}).Preload("Wallet").Take(&user, "id=?", "3")
	assert.Nil(t, result.Error)

	assert.Equal(t, "3", user.Id)
	assert.Equal(t, "3", user.Wallet.Id)
}

func TestRetrieveRelationUserWalletJoin(t *testing.T) {
	var user User
	result := Db().Model(&User{}).Joins("Wallet").Take(&user, "users.id=?", "3")
	assert.Nil(t, result.Error)

	assert.Equal(t, "3", user.Id)
	assert.Equal(t, "1", user.Wallet.Id)
}

func TestAutoCreateUpdateUserWallet(t *testing.T) {
	user := User{
		Id:       "20",
		Password: "rahasia20",
		Name: Name{
			FirstName: "Jumadi 20",
		},
		Wallet: Wallet{
			Id:      "2",
			UserId:  "20",
			Balance: 5_000_000,
		},
	}

	err := Db().Create(&user).Error
	assert.Nil(t, err)
}
