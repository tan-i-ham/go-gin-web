package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDtebase() {
	dsn := "root:123456@tcp(127.0.0.1:3307)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	// database, err := gorm.Open("sqlite3", "test.db")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
