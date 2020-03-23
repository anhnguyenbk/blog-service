package jwtutils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthToken
type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// Response user details
type TokenPayload struct {
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	CreatedAt string   `json:"createdAt"`
	Roles     []string `json:"roles"`
}

// AuthTokenClaim
type AuthTokenClaim struct {
	*jwt.StandardClaims
	TokenPayload
}

const SECRET = "anXvwW"

func ExtractBearerTokenFromRequest(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) < 2 {
		return ""
	}
	return strings.TrimSpace(splitToken[1])
}

func GenerateToken(payload TokenPayload) (AuthToken, error) {
	// Generate token
	expiresAt := time.Now().Add(time.Minute * 60 * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		payload,
	}

	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return AuthToken{}, err
	}

	return AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	}, nil
}

func Verify(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		secret := []byte(SECRET)
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	// Parse and validate token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
