package models

import "time"

type Order struct {
	OrderID     string
	UserID      string
	Status      string
	TotalAmount float64
	CreatedAt   time.Time
}
