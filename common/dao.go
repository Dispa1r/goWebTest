package common

import (
	"github.com/jinzhu/gorm"
	"goWeb/model"
)
var DB *gorm.DB
func GetDB() *gorm.DB{
	return DB
}
func InitDB() *gorm.DB {
	dsn := "root:001228@tcp(127.0.0.1:3306)/ginVueTest?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open("mysql", dsn)
	//
	//if err != nil {
	//	log.Println("data base connect failed")
	//}
	db.LogMode(true)
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
