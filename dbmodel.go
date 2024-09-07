package main

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string          `gorm:"column:first_name"`
	LastName      string          `gorm:"column:last_name"`
	Email         string          `gorm:"column: email"`
	TicketDetails []TicketDetails `gorm:"foreignKey:UserID"`
}

type TrainDetails struct {
	gorm.Model
	TrainName  string       `gorm:"column:train_name"`
	TrainType  string       `gorm:"column:last_name"`
	TrainClass []TrainClass `gorm:"foreignKey:TrainIDReference"`
}

type TrainClass struct {
	gorm.Model
	Class            string         `gorm:"column:train_class"`
	TrainIDReference uint           `gorm:"column:train_id_reference"`
	BookingSeats     []BookingSeats `gorm:"foreignKey:TrainClassReference"`
}

type BookingSeats struct {
	gorm.Model
	SeatNumber          int  `gorm:"column:seat_number"`
	IsBooked            bool `gorm:"column:is_booked"`
	TrainClassReference uint `gorm:"column:train_class_reference"`
}

type PaymentDetails struct {
	gorm.Model
	Amount int    `json:"amount"`
	Status string `json:"status"`
}

type FairDetails struct {
	gorm.Model
	TrainID    int    `json:"train_id"`
	TrainClass string `json:"train_class"`
	Amount     int    `json:"amount"`
}

type TicketDetails struct {
	gorm.Model
	From          string `gorm:"column:from"`
	To            string `gorm:"column:to"`
	Status        string `gorm:"column:ticket_status"`
	Price         int    `gorm:"column:tickets_price"`
	UserID        uint   `gorm:"column:user_id"`
	TransactionID int    `gorm:"column:transaction_id"`
	TrainId       uint   `gorm:"column:train_id"`
	TrainClass    string `gorm:"column:train_class"`
}

func initModel() {
	fmt.Println("Initialize Model---")
	db.AutoMigrate(&User{})
	db.AutoMigrate(&TrainDetails{})
	db.AutoMigrate(&TrainClass{})
	db.AutoMigrate(&BookingSeats{})
	db.AutoMigrate(&PaymentDetails{})
	db.AutoMigrate(&FairDetails{})
	db.AutoMigrate(&TicketDetails{})
	fmt.Println("Model Initialized")
}
