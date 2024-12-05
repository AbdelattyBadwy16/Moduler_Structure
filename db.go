package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error

	dsn := "host=localhost user=postgres password=01142912894 dbname=GoTest port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(&Brand{}, &Product{}, &Tag{})

	if err != nil {
		fmt.Println("Faild To Migrate. ", err)
	}
}
