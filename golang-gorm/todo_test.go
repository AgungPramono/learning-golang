package golang_gorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteTodo(t *testing.T) {
	todo := Todo{
		UserId:      "3",
		Title:       "Todo 3",
		Description: "Description 3",
	}

	err := Db().Create(&todo).Error
	assert.Nil(t, err)

	result := Db().Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	result = Db().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))

	var deleteTodo Todo
	result = Db().Take(&deleteTodo, "id=?", "3")
	assert.NotNil(t, result.Error)
}
