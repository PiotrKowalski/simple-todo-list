package user

import (
	"context"
	"errors"
	"simple-todo-list/internal/domain"
	"simple-todo-list/internal/domain/auth"
	"simple-todo-list/internal/domain/hashing"
	"simple-todo-list/internal/domain/user"
	userDto "simple-todo-list/internal/dtos/user"
)

type RegisterHandler struct {
	Repo interface {
		domain.GenericSaveRepo[user.User]
		domain.GenericGetForMatchingRepo[user.User]
	}
	HashService hashing.HashService
	AuthService auth.Service
}

func (h *RegisterHandler) Handle(ctx context.Context, input userDto.RegisterInput) (userDto.RegisterOutput, error) {
	matching, err := h.Repo.FindForMatching(ctx, user.HasUsernameSpecification{Username: input.Username})
	if err != nil {
		return userDto.RegisterOutput{}, err
	}
	if matching != nil {
		return userDto.RegisterOutput{}, errors.New("user already exists")
	}

	hashedPassword, err := h.HashService.Hash(input.Password)
	if err != nil {
		return userDto.RegisterOutput{}, err
	}

	createdUser, err := h.Repo.Save(ctx, user.NewUser(input.Username, hashedPassword, input.Email))
	if err != nil {
		return userDto.RegisterOutput{}, err
	}

	token, err := h.AuthService.GetJwtToken(*createdUser)
	if err != nil {
		return userDto.RegisterOutput{}, err
	}

	return userDto.NewRegisterOutput(createdUser.Id, token), nil
}
