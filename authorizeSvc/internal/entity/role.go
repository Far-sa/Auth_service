package entity

type Role struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
}

// UserRole represents the many-to-many relationship between users and roles.
type UserRole struct {
	UserID string `json:"user_id" gorm:"primaryKey"`
	RoleID string `json:"role_id" gorm:"primaryKey"`
}
