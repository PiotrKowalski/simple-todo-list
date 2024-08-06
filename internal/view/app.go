package view

import (
	"context"
	"simple-todo-list/internal/dtos"
)

type app interface {
	CreateTodoList(ctx context.Context, in dtos.CreateTodoListInput) (dtos.CreateTodoListOutput, error)
	GetByIdTodoList(ctx context.Context, in dtos.GetByIdTodoListInput) (dtos.GetByIdTodoListOutput, error)
	Register(ctx context.Context, in dtos.RegisterInput) (dtos.RegisterOutput, error)
	Login(ctx context.Context, in dtos.LoginInput) (dtos.LoginOutput, error)
}
