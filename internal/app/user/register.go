package user

import (
	"context"
	"simple-todo-list/internal/domain"
	"simple-todo-list/internal/domain/auth"
	"simple-todo-list/internal/domain/hashing"
	"simple-todo-list/internal/domain/user"
	userDto "simple-todo-list/internal/dtos/user"
)

type RegisterHandler struct {
	Repo        domain.GenericSaveRepo[user.User]
	HashService hashing.HashService
	AuthService auth.Service
}

func (h *RegisterHandler) Handle(ctx context.Context, input userDto.RegisterInput) (userDto.RegisterOutput, error) {
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
