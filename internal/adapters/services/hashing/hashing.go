package hashing

import "golang.org/x/crypto/bcrypt"

type Service struct{}

func (s Service) Hash(toHash string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(toHash), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(password), nil
}

func (s Service) CompareHashAndClear(hashed string, clearString string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(clearString))
}

func New() Service {
	return Service{}
}
