package internal

// AuthJson represents a request sent to /users/register and /users/login
type AuthJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewPasswordJson represents a request sent to /users/new-password
type NewPasswordJson struct {
	NewPassword string `json:"newPassword"`
}
