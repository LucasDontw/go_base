package model

import "time"

type Account struct {
	ID         int64     `gorm:"column:id;primary_key"`
	UserID     string    `gorm:"column:user_id"`
	Password   string    `gorm:"column:password"`
	Nickname   string    `gorm:"column:nickname"`
	Created_at time.Time `gorm:"column:created_at"`
	Updated_at time.Time `gorm:"column:updated_at"`
}

func (a Account) TableName() string { //指定此結構對應到哪個table
	return "cms_account.accounts"
}