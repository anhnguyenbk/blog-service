package user

import (
	"fmt"
	"net/http"

	"github.com/anhnguyenbk/blog-service/internal/util/jwtutils"
	"github.com/anhnguyenbk/blog-service/internal/util/requestutils"
	"github.com/anhnguyenbk/blog-service/internal/util/responseutils"
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

	payload := jwtutils.TokenPayload{user.Username, user.Email, user.CreatedAt, user.Roles}
	authToken, err := jwtutils.GenerateToken(payload)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}

	responseutils.ResponseJSON(w, authToken)
}
