package entity

type Permission struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
}

// RolePermission represents the many-to-many relationship between roles and permissions.
type RolePermission struct {
	RoleID       int `json:"role_id" gorm:"primaryKey"`
	PermissionID int `json:"permission_id" gorm:"primaryKey"`
}
