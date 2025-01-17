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
