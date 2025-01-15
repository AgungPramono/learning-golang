package golang_gorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	connection := OpenConnection()
	assert.NotNil(t, connection)
}

func TestExecuteRawSql(t *testing.T) {
	err := OpenConnection().Exec("insert into sample(id, name) value (?,?)", "1", "agung").Error
	assert.Nil(t, err)

	err = OpenConnection().Exec("insert into sample(id, name) value (?,?)", "2", "joko").Error
	assert.Nil(t, err)

	err = OpenConnection().Exec("insert into sample(id, name) value (?,?)", "3", "ahmad").Error
	assert.Nil(t, err)

	err = OpenConnection().Exec("insert into sample(id, name) value (?,?)", "4", "budi").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestRawSQl(t *testing.T) {
	var sample Sample
	err := OpenConnection().Raw("select id,name from sample where id=?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "agung", sample.Name)

	var samples []Sample
	err = OpenConnection().Raw("select id,name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}
