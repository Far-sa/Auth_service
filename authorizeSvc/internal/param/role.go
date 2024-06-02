package param

// UserEvent represents an event from the UserService
type UserEvent struct {
	Type string // The type of the event (e.g., "USER_CREATED", "USER_UPDATED")
	User User   // The user that the event pertains to
}

// User represents a user in the UserService
type User struct {
	ID string // The ID of the user
	// ... other fields ...
}

type RoleAssignmentRequest struct {
	UserID string `json:"userId"`
}

type RoleAssignmentResponse struct {
	Message string `json:"message"`
}
type RoleUpdateRequest struct {
	UserID  string `json:"userId"`
	NewRole string `json:"newRole"` // Assuming role is represented as a string
}
