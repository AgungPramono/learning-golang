package golang_gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-gorm/model"
	"gorm.io/gorm/clause"
	"math"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	wallet := model.Wallet{
		Id:      "1",
		UserId:  "3",
		Balance: 40_000_000,
	}
	err := Db().Create(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRelationUserWallet(t *testing.T) {
	var user model.User
	result := Db().Model(&model.User{}).Preload("Wallet").Take(&user, "id=?", "3")
	assert.Nil(t, result.Error)

	assert.Equal(t, "3", user.Id)
	assert.Equal(t, "3", user.Wallet.Id)
}

func TestRetrieveRelationUserWalletJoin(t *testing.T) {
	var user model.User
	result := Db().Model(&model.User{}).Joins("Wallet").Take(&user, "users.id=?", "3")
	assert.Nil(t, result.Error)

	assert.Equal(t, "3", user.Id)
	assert.Equal(t, "1", user.Wallet.Id)
}

func TestAutoCreateUpdateUserWallet(t *testing.T) {
	user := model.User{
		Id:       "20",
		Password: "rahasia20",
		Name: model.Name{
			FirstName: "Jumadi 20",
		},
		Wallet: model.Wallet{
			Id:      "2",
			UserId:  "20",
			Balance: 5_000_000,
		},
	}

	err := Db().Create(&user).Error
	assert.Nil(t, err)
}
func TestSkipAutoCreateUpdateUserWallet(t *testing.T) {
	user := model.User{
		Id:       "21",
		Password: "rahasia21",
		Name: model.Name{
			FirstName: "Suhadi 21",
		},
		Wallet: model.Wallet{
			Id:      "3",
			UserId:  "21",
			Balance: 2_000_000,
		},
	}

	err := Db().Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}

// ambil wallet, termasuk user pemilik wallet, sekaligus addressnya
func TestPreloadingNested(t *testing.T) {
	var wallet model.Wallet
	err := Db().Preload("User.Addresses").Take(&wallet, "id=?", "3").Error
	assert.Nil(t, err)

	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}

func TestPreloadAll(t *testing.T) {
	var user model.User
	err := Db().Preload(clause.Associations).Take(&user, "id=?", "3").Error
	assert.Nil(t, err)
}

func TestJoinQuery(t *testing.T) {
	var user []model.User
	err := Db().Joins("join wallets on wallets.user_id = users.id").Find(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(user))

	user = []model.User{}
	err = Db().Joins("Wallet").Find(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(user))
}

func TestJoinWithCondition(t *testing.T) {
	var user []model.User
	err := Db().Joins("join wallets on wallets.id = users.id where wallets.balance >?", 500000).Find(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(user))

	user = []model.User{}
	err = Db().Joins("Wallet").Where("Wallet.balance >?", 500000).Find(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(user))
}

func TestAggregation(t *testing.T) {
	var count int64
	err := Db().Model(&model.User{}).Joins("Wallet").Where("Wallet.balance>?", 500000).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(3), count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregationManual(t *testing.T) {
	var result AggregationResult
	err := Db().Model(&model.Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance",
		"max(balance) as max_balance", "avg(balance) as avg_balance").Take(&result).Error
	assert.Nil(t, err)

	fmt.Println("total balance:", result.TotalBalance)
	fmt.Println("total balance:", result.MinBalance)
	fmt.Println("total balance:", result.MaxBalance)
	fmt.Println("total balance:", math.Round(result.AvgBalance))
}
