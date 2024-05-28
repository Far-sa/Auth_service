package param

// * Paramas

type UserInfo struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	CreateAt    string `json:"create_at"`
}

type ProfileRequest struct {
	UserID string `json:"user_id"`
}
type ProfileResponse struct {
	UserInfo
}
