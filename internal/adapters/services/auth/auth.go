package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"simple-todo-list/internal/domain/user"
	"time"
)

type Role string

type jwtCustomClaims struct {
	Username string `json:"username"`
	Role     Role   `json:"role"`
	jwt.RegisteredClaims
}

type Service struct{}

func (s Service) GetJwtToken(user user.User) (string, error) {
	claims := &jwtCustomClaims{
		user.Username,
		"regular",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	file, err := os.ReadFile("./pkg/auth/private.pem")
	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(file)
	if err != nil {
		return "", err
	}

	// Generate encoded token and send it as response.
	t, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return t, nil

}

func New() Service {
	return Service{}
}
