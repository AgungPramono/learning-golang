package golang_gorm

import "time"

type User struct {
	Id          string    `gorm:"primary_key;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"`
}

// override nama table user
func (u *User) TableName() string {
	return "users"
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

type UserLog struct {
	ID        string `gorm:"primary_key;column:id;autoIncrement"`
	UserId    string `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *UserLog) TableName() string {
	return "user_log"
}
