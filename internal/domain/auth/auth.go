package auth

import "simple-todo-list/internal/domain/user"

type Service interface {
	GetJwtToken(user user.User) (string, error)
}
