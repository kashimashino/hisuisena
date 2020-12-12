package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var db *gorm.DB
var err error
func InitDb()  {
	db , err = gorm.Open("mysql","username:password@(localhost)/hackweek?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{},&Clothes{})
}

//布置全局db使用
func GetDB() *gorm.DB {
	return db
}