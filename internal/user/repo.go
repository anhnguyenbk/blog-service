package user

import (
	"github.com/anhnguyenbk/blog-service/internal/util/dynamodbutils"
	"golang.org/x/crypto/bcrypt"
)

var tableName string = "blog_users"

func Authentication(request LoginRequest) (User, error) {
	var user = User{}
	err := dynamodbutils.FindByID(tableName, request.Email, &user)
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return User{}, err
	}
	return user, nil
}