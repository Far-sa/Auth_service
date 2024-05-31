package param

type LoginRequest struct {
	Email    string `json:"username_or_email"`
	Password string `json:"password"`
}

// LoginResponse represents the response after a successful login
type LoginResponse struct {
	UserID string `json:"user_id"`
	Tokens Token  `json:"token_pair"`
}

type GetUserResponse struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

// type RegisterRequest struct {
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// // RegisterResponse represents the response after a successful registration
// type RegisterResponse struct {
// 	UserID    int       `json:"user_id"`
// 	Username  string    `json:"username"`
// 	Email     string    `json:"email"`
// 	CreatedAt time.Time `json:"created_at"`
// }

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
