package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/anhnguyenbk/blog-service/internal/auth"
	"github.com/anhnguyenbk/blog-service/internal/util/responseutils"
	"github.com/dgrijalva/jwt-go"
)

// CORS middleware
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// CORS middleware
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("RequestPath: " + r.URL.Path + ", RequestMethod: " + r.Method)
		t := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("Execution time: %s \n", time.Now().Sub(t).String())

	})
}

// Authentication Middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO Except URLs config
		if isAuthExceptURI(r.RequestURI) {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := extractBearerTokenFromRequest(r)
		if tokenString == "" {
			responseutils.ResponseErrorWithStatus(w, 401, fmt.Errorf("Invalid or missing authorization token"))
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			secret := []byte("anXvwW")
			return secret, nil
		})

		if err != nil {
			fmt.Println("Authenticate error: " + err.Error())
			responseutils.ResponseErrorWithStatus(w, 401, fmt.Errorf("Invalid or missing authorization token"))
			return
		}

		// Parse and validate token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)

			roles := auth.ExtractRoles(claims)
			if !auth.HasRole(roles, "ROLE_ADMIN") && !auth.HasRole(roles, "ROLE_SYSTEM") {
				responseutils.ResponseErrorWithStatus(w, 401, fmt.Errorf("Unauthorized"))
				return
			}
		} else {
			responseutils.ResponseErrorWithStatus(w, 401, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isAuthExceptURI(uri string) bool {
	return uri == "/auth/token"
}

func extractBearerTokenFromRequest(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) < 2 {
		return ""
	}
	return strings.TrimSpace(splitToken[1])
}
