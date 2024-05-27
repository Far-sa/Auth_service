package param

// * Paramas

type UserInfo struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CreateAt    string `json:"create_at"`
}

type ProfileRequest struct {
	UserID string `json:"user_id"`
}
type ProfileResponse struct {
	Name string `json:"name"`
}
