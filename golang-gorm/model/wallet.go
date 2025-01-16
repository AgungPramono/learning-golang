package model

import "time"

type Wallet struct {
	Id      string `gorm:"primary_key;column:id"`
	UserId  string `gorm:"column:user_id"`
	Balance int64  `gorm:"column:balance"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	User *User `gorm:"foreignkey:user_id;references:id"`
}

func (w *Wallet) TableName() string {
	return "wallets"
}
