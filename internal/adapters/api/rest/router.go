package rest

import (
	"context"
	"github.com/labstack/echo/v4"
	"simple-todo-list/internal/dtos"
)

type app interface {
	CreateTodoList(ctx context.Context, in dtos.CreateTodoListInput) (dtos.CreateTodoListOutput, error)
	GetByIdTodoList(ctx context.Context, in dtos.GetByIdTodoListInput) (dtos.GetByIdTodoListOutput, error)
}

func New(e *echo.Group, app app) {
	e.POST("/todolist", createTodoList(app))
	e.GET("/todolist/:id", getByIdTodoList(app))
}
