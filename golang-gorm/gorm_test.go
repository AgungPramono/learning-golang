package golang_gorm

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"strconv"
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

func TestSqlRows(t *testing.T) {
	rows, err := OpenConnection().Raw("select id,name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()
	var samples []Sample

	for rows.Next() {
		var id string
		var name string

		err = rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}
	assert.Equal(t, 4, len(samples))
}

func TestScanRows(t *testing.T) {
	rows, err := OpenConnection().Raw("select id,name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()
	var samples []Sample

	for rows.Next() {
		err := OpenConnection().ScanRows(rows, &samples)
		assert.Nil(t, err)
	}
	assert.Equal(t, 4, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := &User{
		Id:       "1",
		Password: "12345",
		Name: Name{
			FirstName:  "Budi",
			MiddleName: "Siregar",
			LastName:   "Kapal Laut",
		},
		Information: "ini akan di ignore",
	}
	response := OpenConnection().Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			Id:       strconv.Itoa(i),
			Password: "1234" + strconv.Itoa(i),
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
		})
	}

	result := OpenConnection().Create(&users)
	assert.NotNil(t, result)
	assert.Equal(t, int64(8), result.RowsAffected)
}

func TestTransaction(t *testing.T) {
	err := OpenConnection().Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			Id:       "10",
			Password: "1234",
			Name: Name{
				FirstName: "User " + strconv.Itoa(10),
			},
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			Id:       "11",
			Password: "1234",
			Name: Name{
				FirstName: "User " + strconv.Itoa(11),
			},
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			Id:       "12",
			Password: "1234",
			Name: Name{
				FirstName: "User " + strconv.Itoa(12),
			},
		}).Error
		if err != nil {
			return err
		}

		return nil
	})
	assert.Nil(t, err)
}

func TestTransactionErrorRollback(t *testing.T) {
	err := OpenConnection().Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			Id:       "13",
			Password: "1234",
			Name: Name{
				FirstName: "User " + strconv.Itoa(13),
			},
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			Id:       "11",
			Password: "1234",
			Name: Name{
				FirstName: "User " + strconv.Itoa(14),
			},
		}).Error
		if err != nil {
			return err
		}
		return nil
	})
	assert.NotNil(t, err)
}

func TestManualTransaction(t *testing.T) {
	tx := OpenConnection().Begin()
	defer tx.Rollback()

	err := tx.Create(&User{
		Id:       "13",
		Password: "1234",
		Name: Name{
			FirstName: "User " + strconv.Itoa(13),
		},
	}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{
		Id:       "14",
		Password: "1234",
		Name: Name{
			FirstName: "User " + strconv.Itoa(14),
		},
	}).Error
	assert.Nil(t, err)

	//jika tidak ada error maka lakukan commit
	if err == nil {
		tx.Commit()
	}
}

func TestManualTransactionError(t *testing.T) {
	tx := OpenConnection().Begin()
	defer tx.Rollback()

	err := tx.Create(&User{
		Id:       "15",
		Password: "1234",
		Name: Name{
			FirstName: "User " + strconv.Itoa(15),
		},
	}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{
		Id:       "14",
		Password: "1234",
		Name: Name{
			FirstName: "User " + strconv.Itoa(14),
		},
	}).Error
	assert.Nil(t, err)

	//jika tidak ada error maka lakukan commit
	if err == nil {
		tx.Commit()
	}
}

func TestQuerySingleObject(t *testing.T) {
	user := &User{}
	err := OpenConnection().First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.Id)

	user = &User{}
	err = OpenConnection().Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.Id)
}

func TestQuerySingleObjectInlineCondition(t *testing.T) {
	user := &User{}
	err := OpenConnection().Take(&user, "id=?", "6").Error
	assert.Nil(t, err)
	assert.Equal(t, "6", user.Id)
	assert.Equal(t, "User 6", user.Name.FirstName)
}

func TestQueryAllObject(t *testing.T) {
	var users []User
	err := OpenConnection().Find(&users, "id in ?", []string{"1", "2", "3", "4"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}
