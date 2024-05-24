package models

import "time"

type User struct {
	UserID    string
	Name      string
	Email     string
	CreatedAt time.Time
}
