package entities

// User represents a user entity
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// ... other user fields
	PasswordHash []byte `json:"-"`                       // Password hash (optional)
	AccessToken  string `json:"access_token,omitempty"`  // Access token for the user (optional)
	RefreshToken string `json:"refresh_token,omitempty"` // Refresh token for the user (optional)
}
