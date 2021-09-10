package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const connStr string = "host=localhost user=postgres password=Weight!123 dbname=todolist port=5432 sslmode=disable"

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&Task{})

	DB = database

}
