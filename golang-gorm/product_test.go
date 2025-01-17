package golang_gorm

import (
	"github.com/stretchr/testify/assert"
	"golang-gorm/model"
	"testing"
)

func TestCreateManyToMany(t *testing.T) {
	product := model.Product{
		ID:    "P001",
		Name:  "Product 001",
		Price: 1000000,
	}

	err := Db().Create(&product).Error
	assert.Nil(t, err)

	err = Db().Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "21",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)

	err = Db().Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "22",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToManyProduct(t *testing.T) {
	var product model.Product
	err := Db().Preload("LikedByUsers").Take(&product, "id=?", "P001").Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))
}

func TestPreloadManyToManyUser(t *testing.T) {
	var user model.User
	err := Db().Preload("LikeProducts").Take(&user, "id=?", "21").Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(user.LikeProducts))
}

func TestAssociationFind(t *testing.T) {
	//select produk dulu
	var product model.Product
	err := Db().Take(&product, "id=?", "P001").Error
	assert.Nil(t, err)

	//cari User yang menyukai produk P001 dengan first name like 'ahmad'
	var user []model.User
	err = Db().Model(&product).Where("users.first_name LIKE ?", "ahmad%").Association("LikedByUsers").Find(&user)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(user))
}
