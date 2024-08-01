package app

import (
	"context"
	"simple-todo-list/internal/dtos"
	"simple-todo-list/internal/dtos/user"
)

type RestApp interface {
	CreateTodoList(ctx context.Context, in dtos.CreateTodoListInput) (dtos.CreateTodoListOutput, error)
	GetByIdTodoList(ctx context.Context, in dtos.GetByIdTodoListInput) (dtos.GetByIdTodoListOutput, error)
	Register(ctx context.Context, in user.RegisterInput) (user.RegisterOutput, error)
}
