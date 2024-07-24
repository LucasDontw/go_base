package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID         int64     `gorm:"column:id;primary_key"`
	UserID     string    `gorm:"column:user_id"`
	Password   string    `gorm:"column:password"`
	Nickname   string    `gorm:"column:nickname"`
	Created_at time.Time `gorm:"column:created_at"`
	Updated_at time.Time `gorm:"column:updated_at"`
}

// func (a Account) TableName() string { //指定此結構對應到哪個table
// 	return "accounts"
// }

func main() {
	db := connDB()
	var accounts []Account
	if err := db.Table("accounts").Find(&accounts).Error; err != nil { //會直接去映射到struct + s 的表 = accounts，或者用Table()直接指定跟TableName一樣用法
		return
	}

	fmt.Println(accounts)

	db.Where("id=?", 2).First(&accounts)

	fmt.Println(accounts)
}

func connDB() *gorm.DB {
	mysqlDB, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/cms_account?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic(err)
	}

	db, err := mysqlDB.DB()

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return mysqlDB
}
