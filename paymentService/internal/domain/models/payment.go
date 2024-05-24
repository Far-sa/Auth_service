package models

import "time"

type Payment struct {
	PaymentID string
	OrderID   string
	Amount    float64
	Method    string
	CreatedAt time.Time
}
