package user

import (
	"fmt"
	"net/http"

	"time"

	"github.com/anhnguyenbk/blog-service/internal/util/requestutils"
	"github.com/anhnguyenbk/blog-service/internal/util/responseutils"
	"github.com/dgrijalva/jwt-go"
)

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest

	err := requestutils.ParseJSONBody(r, &loginRequest)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}

	// Authenticate for user
	user, err := Authentication(loginRequest)
	if err != nil {
		// Log the real error
		fmt.Println(err)

		responseutils.ResponseErrorWithStatus(w, 401, fmt.Errorf("Invalid username or password"))
		return
	}

	// Generate token
	userDetails := UserDetails{user.Username, user.Email, user.CreatedAt, user.Roles}
	expiresAt := time.Now().Add(time.Minute * 60 * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		userDetails,
	}

	tokenString, error := token.SignedString([]byte("anXvwW"))
	if error != nil {
		responseutils.ResponseError(w, err)
		return
	}

	authToken := AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	}
	responseutils.ResponseJSON(w, authToken)
}
