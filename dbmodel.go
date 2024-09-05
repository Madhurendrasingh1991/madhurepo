package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column: email"`
}

type TrainDetails struct {
	gorm.Model
	TrainName  string       `gorm:"column:train_name"`
	TrainType  string       `gorm:"column:last_name"`
	TrainClass []TrainClass `gorm:"column:train_class"`
}

type TrainClass struct {
	gorm.Model
	Class        string         `gorm:"column:train_class"`
	BookingSeats []BookingSeats `gorm:"column:booking_seats"`
	TrainDetails []TrainDetails
}

type BookingSeats struct {
	gorm.Model
	SeatNumber int  `gorm:"column:seat_number"`
	IsBooked   bool `gorm:"column:is_booked"`
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
	User          User   `gorm:"column:user_details"`
	TransactionID int    `gorm:"column:transaction_id"`
	TrainId       uint   `gorm:"column:train_id"`
	TrainClass    string `gorm:"column:train_class"`
}

func initModel() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&TrainDetails{})
	db.AutoMigrate(&TrainClass{})
	db.AutoMigrate(&BookingSeats{})
	db.AutoMigrate(&PaymentDetails{})
	db.AutoMigrate(&FairDetails{})
	db.AutoMigrate(&TicketDetails{})
}
