package dtos

import "time"

type CreateBookingReq struct {
	CarsID     string
	UserID     string
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice float64
	Status
	PaymentMethod
}
