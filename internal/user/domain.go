package user

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// System user details
type User struct {
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	CreatedAt string   `json:"createdAt"`
	Roles     []string `json:"roles"`
}
