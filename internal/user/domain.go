package user

import "github.com/dgrijalva/jwt-go"

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

// Response user details
type UserDetails struct {
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	CreatedAt string   `json:"createdAt"`
	Roles     []string `json:"roles"`
}

// AuthToken
type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// AuthTokenClaim
type AuthTokenClaim struct {
	*jwt.StandardClaims
	UserDetails
}
