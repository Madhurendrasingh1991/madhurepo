package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func initializeDbConnection() {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=root dbname=train_ticket port=5432 sslmode=disable TimeZone=Asia/Kolkata",
		PreferSampleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
}

func main() {

	initModel()
}
