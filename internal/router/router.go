package router

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"simple-todo-list/internal/dtos"
)

type app interface {
	CreateTodoList(ctx context.Context, in dtos.CreateTodoListInput) (dtos.CreateTodoListOutput, error)
	GetByIdTodoList(ctx context.Context, in dtos.GetByIdTodoListInput) (dtos.GetByIdTodoListOutput, error)
}

func New(app app) (*echo.Echo, error) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	v1 := e.Group("/v1")
	v1.POST("/todolist", createTodoList(app))
	v1.GET("/todolist/:id", getByIdTodoList(app))

	return e, nil
}
