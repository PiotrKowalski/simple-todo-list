package user

import (
	"context"
	"errors"
	"simple-todo-list/internal/domain"
	"simple-todo-list/internal/domain/auth"
	"simple-todo-list/internal/domain/hashing"
	"simple-todo-list/internal/domain/user"
	"simple-todo-list/internal/dtos"
)

type LoginHandler struct {
	Repo interface {
		domain.GenericSaveRepo[user.User]
		domain.GenericGetForMatchingRepo[user.User]
	}
	HashService hashing.HashCheckService
	AuthService auth.Service
}

func (h *LoginHandler) Handle(ctx context.Context, input dtos.LoginInput) (dtos.LoginOutput, error) {
	matching, err := h.Repo.FindForMatching(ctx, user.HasUsernameSpecification{Username: input.Username})
	if err != nil {
		return dtos.LoginOutput{}, err
	}
	if len(matching) == 0 {
		return dtos.LoginOutput{}, errors.New("wrong credentials")
	}
	userFromDb := matching[0]
	err = h.HashService.CompareHashAndClear(userFromDb.Password, input.Password)
	if err != nil {
		return dtos.LoginOutput{}, err
	}

	token, err := h.AuthService.GetJwtToken(userFromDb)
	if err != nil {
		return dtos.LoginOutput{}, err
	}

	return dtos.LoginOutput{JWT: token}, nil
}
