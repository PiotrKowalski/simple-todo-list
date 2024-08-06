package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"simple-todo-list/internal/domain/user"
	"simple-todo-list/pkg/auth"
)

type Service struct{}

func (s Service) GetJwtToken(user user.User) (string, error) {
	return auth.GenerateJwt(auth.JwtCustomClaims{
		Username:         user.Username,
		Role:             "regular",
		RegisteredClaims: jwt.RegisteredClaims{},
	})
}

func New() Service {
	return Service{}
}
