package models

import "time"

type Inventory struct {
	ProductID   string
	Quantity    int32
	LastUpdated time.Time
}
