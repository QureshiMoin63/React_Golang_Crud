package models

import (
	"fmt"
	"restapi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:@tcp(127.0.0.1:3306)/phpmyadmin?parseTime=true"

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot Connect to the DB")
	}
	DB.AutoMigrate(&models.User{}, &models.Author{}, &models.Book{})
}

func GetDB() *gorm.DB {
	return DB
}
