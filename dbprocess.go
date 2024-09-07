package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func stablishDbConnection() {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=root dbname=trainticket port=5432 sslmode=disable TimeZone=Asia/Kolkata",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
}
