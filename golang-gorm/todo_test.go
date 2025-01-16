package golang_gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-gorm/model"
	"testing"
)

func TestSoftDelete(t *testing.T) {
	todo := model.Todo{
		UserId:      "6",
		Title:       "Todo 6",
		Description: "Description 6",
	}

	err := Db().Create(&todo).Error
	assert.Nil(t, err)

	result := Db().Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.DeletedAt)

	var todos []model.Todo
	result = Db().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))

	var deleteTodo model.Todo
	result = Db().Take(&deleteTodo, "id=?", "6")
	assert.NotNil(t, result.Error)
}

/*
*
query dengan mengabaikan data yang sudah di hapus dengan metode Soft Delete
*/
func TestUnscoped(t *testing.T) {
	var todo model.Todo
	err := Db().Unscoped().First(&todo, "id=?", "7").Error
	assert.Nil(t, err)
	fmt.Println(todo)

	err = Db().Unscoped().Delete(&todo).Error
	assert.Nil(t, err)

	var todos []model.Todo
	err = Db().Unscoped().Find(&todos).Error
	assert.Nil(t, err)
}
