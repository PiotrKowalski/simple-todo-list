package user

import "github.com/google/uuid"

type Option func(user *User)

func NewUser(
	username, password, email string,
	opts ...Option,
) *User {
	u := &User{
		Id:       uuid.New().String(),
		Username: username,
		Password: password,
		Email:    email,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u

}
