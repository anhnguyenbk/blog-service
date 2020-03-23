package auth

import (
	"github.com/anhnguyenbk/blog-service/internal/util/strutils"
	"github.com/dgrijalva/jwt-go"
)

func HasRole(roles []string, role string) bool {
	return strutils.Contains(roles, role)
}

func ExtractRoles(claims jwt.MapClaims) []string {
	roleInterfaces := claims["roles"].([]interface{})
	roles := make([]string, len(roleInterfaces))
	for i, v := range roleInterfaces {
		roles[i] = v.(string)
	}
	return roles
}
