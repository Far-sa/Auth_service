package param

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}
type LoginResponse struct {
	UserID int      `json:"user_id"`
	Tokens []string `json:"tokens"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// !
type UserResponse struct {
	ID           int    `json:"id"`
	PasswordHash string `json:"password_hash"`
	Error        string `json:"error,omitempty"`
}

type RegisterUserRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type RegisterUserResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
