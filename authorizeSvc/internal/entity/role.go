package entity

type Role struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
}

// UserRole represents the many-to-many relationship between users and roles.
type UserRole struct {
	UserID int `json:"user_id" gorm:"primaryKey"`
	RoleID int `json:"role_id" gorm:"primaryKey"`
}
