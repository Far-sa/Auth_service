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
