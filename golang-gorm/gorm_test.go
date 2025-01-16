package golang_gorm

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	connection := Db()
	assert.NotNil(t, connection)
}

func TestExecuteRawSql(t *testing.T) {
	err := Db().Exec("insert into sample(id, name) value (?,?)", "1", "agung").Error
	assert.Nil(t, err)

	err = Db().Exec("insert into sample(id, name) value (?,?)", "2", "joko").Error
	assert.Nil(t, err)

	err = Db().Exec("insert into sample(id, name) value (?,?)", "3", "ahmad").Error
	assert.Nil(t, err)

	err = Db().Exec("insert into sample(id, name) value (?,?)", "4", "budi").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestRawSQl(t *testing.T) {
	var sample Sample
	err := Db().Raw("select id,name from sample where id=?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "agung", sample.Name)

	var samples []Sample
	err = Db().Raw("select id,name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestSqlRows(t *testing.T) {
	rows, err := Db().Raw("select id,name from sample").Rows()
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
	rows, err := Db().Raw("select id,name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()
	var samples []Sample

	for rows.Next() {
		err := Db().ScanRows(rows, &samples)
		assert.Nil(t, err)
	}
	assert.Equal(t, 4, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := &User{
		Id:       "15",
		Password: "rahasia",
		Name: Name{
			FirstName:  "Joko",
			MiddleName: "",
			LastName:   "",
		},
		Information: "ini akan di ignore",
	}
	response := Db().Create(&user)
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

	result := Db().Create(&users)
	assert.NotNil(t, result)
	assert.Equal(t, int64(8), result.RowsAffected)
}

func TestTransaction(t *testing.T) {
	err := Db().Transaction(func(tx *gorm.DB) error {
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
	err := Db().Transaction(func(tx *gorm.DB) error {
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
	tx := Db().Begin()
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
	tx := Db().Begin()
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
	err := Db().First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.Id)

	user = &User{}
	err = Db().Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.Id)
}

func TestQuerySingleObjectInlineCondition(t *testing.T) {
	user := &User{}
	err := Db().Take(&user, "id=?", "6").Error
	assert.Nil(t, err)
	assert.Equal(t, "6", user.Id)
	assert.Equal(t, "User 6", user.Name.FirstName)
}

func TestQueryAllObject(t *testing.T) {
	var users []User
	err := Db().Find(&users, "id in ?", []string{"1", "2", "3", "4"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestQueryCondition(t *testing.T) {
	var users []User
	err := Db().
		Where("first_name like ?", "%User%").
		Where("password=?", "1234").
		Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))
}

func TestQueryOrOperator(t *testing.T) {
	var users []User
	err := Db().
		Where("first_name like ?", "%User%").
		Or("password=?", "12345").
		Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))
}

func TestQueryNotOperator(t *testing.T) {
	var users []User
	err := Db().
		Not("first_name like ?", "%User%").
		Where("password=?", "12345").
		Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestSelectField(t *testing.T) {
	var users []User
	err := Db().Select("id", "first_name").Find(&users).Error
	assert.Nil(t, err)
	for _, user := range users {
		assert.NotNil(t, user)
		assert.NotEqual(t, "", user.Id)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	assert.Equal(t, 14, len(users))
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User 10",
			LastName:  "", // tidak bisa harus menggunakan map condition
		},
		Password: "1234",
	}

	var users []User
	err := Db().Where(&userCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestMapCondition(t *testing.T) {
	mapCondition := map[string]interface{}{
		"middle_name": "",
		"last_name":   "",
	}

	var users []User
	err := Db().Where(mapCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))
}

func TestOrderLimitAndOffset(t *testing.T) {
	var users []User
	err := Db().Order("id asc, first_name desc").Limit(5).Offset(5).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))
}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	err := Db().Model(&User{}).Select("id,first_name,last_name").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))
}

func TestUpdate(t *testing.T) {
	user := User{}
	err := Db().Take(&user, "id=?", "15").Error
	assert.Nil(t, err)

	//ubah data
	user.Name.MiddleName = "waluyo"
	user.Name.LastName = "Nugroho"
	user.Password = "rahasiabanget"

	//simpan data
	err = Db().Save(&user).Error
	assert.Nil(t, err)
}

func TestUpdateSelectedColumn(t *testing.T) {
	err := Db().Model(&User{}).Where("id=?", "15").Updates(map[string]interface{}{
		"middle_name": "Purnomo",
		"last_name":   "",
	}).Error
	assert.Nil(t, err)

	err = Db().Model(&User{}).Where("id=?", "15").Update("password", "rahasialagi").Error
	assert.Nil(t, err)

	err = Db().Model(&User{}).Where("id=?", "15").Updates(User{
		Name: Name{
			FirstName: "Budiono",
			LastName:  "Utomo",
		},
	}).Error
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	var user User
	//cari data dulu
	result := Db().Take(&user, "id=?", "4")
	assert.Nil(t, result.Error)

	result = Db().Delete(&user)
	assert.Nil(t, result.Error)

	//tanpa cari dulu
	result = Db().Delete(&User{}, "id=?", "5")
	assert.Nil(t, result.Error)
	//
	////delete dengan where
	result = Db().Where("id=?", "6").Delete(&User{})
	assert.Nil(t, result.Error)
}

func TestLock(t *testing.T) {
	err := Db().Transaction(func(tx *gorm.DB) error {

		var user User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id=?", "3").Error
		if err != nil {
			return err
		}

		user.Name.FirstName = "Joko"
		user.Name.MiddleName = "Anwar"
		err = tx.Save(&user).Error
		return err
	})
	assert.Nil(t, err)
}
