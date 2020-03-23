package middleware

import (
	"fmt"
	"net/http"

	"github.com/anhnguyenbk/blog-service/internal/auth"
	"github.com/anhnguyenbk/blog-service/internal/util/jwtutils"
	"github.com/anhnguyenbk/blog-service/internal/util/responseutils"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO Except URLs config
		if isAuthExceptURI(r.RequestURI) {
			next.ServeHTTP(w, r)
			return
		}

		// Extract token
		tokenString := jwtutils.ExtractBearerTokenFromRequest(r)
		if tokenString == "" {
			responseutils.ResponseErrorWithStatus(w, 401, fmt.Errorf("Invalid or missing authorization token"))
			return
		}

		// Verify token
		claims, err := jwtutils.Verify(tokenString)
		if err != nil {
			fmt.Println("Authenticate error: " + err.Error())
			responseutils.ResponseErrorWithStatus(w, 401, fmt.Errorf("Invalid or missing authorization token"))
			return
		}

		// check roles
		roles := auth.ExtractRoles(claims)
		if !auth.HasRole(roles, "ROLE_ADMIN") && !auth.HasRole(roles, "ROLE_SYSTEM") {
			responseutils.ResponseErrorWithStatus(w, 401, fmt.Errorf("Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isAuthExceptURI(uri string) bool {
	return uri == "/auth/token"
}
