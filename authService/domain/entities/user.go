package entities

// User represents a user entity
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// ... other user fields
	PasswordHash []byte `json:"-"` // Password hash (optional)
}

