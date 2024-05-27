package entity

import "time"

type User struct {
	UserID    string
	Name      string
	Email     string
	CreatedAt time.Time
}
