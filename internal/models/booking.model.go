package models

import (
	"Rental-car/pkg/common"
	"time"

	"gorm.io/gorm"
)

const {
	STATUS_BOOKING_NOT_PAID = "NOT_PAID",
	STATUS_BOOKING_BOOKING = "BOOKING"
}

type Booking struct {
	common.ModelsWithID
	CarsID        string    `json:"cars_id" gorm:"foreignKey:cars_id"`
	UserID        string    `json:"user_id" gorm:"foreignKey:user_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	TotalPrice    float64   `json:"total_price"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
}

func (c *Booking) TableName() string {
	return "Booking"
}
func (c Booking) BeforeCreate(db gorm.DB) error {
	c.GenerateUUID("BKG")
	return nil
}
