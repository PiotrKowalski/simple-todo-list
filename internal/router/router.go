package router

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"simple-todo-list/internal/adapters/api/rest"
	"simple-todo-list/internal/dtos"
	"simple-todo-list/internal/view"
)

type app interface {
	CreateTodoList(ctx context.Context, in dtos.CreateTodoListInput) (dtos.CreateTodoListOutput, error)
	GetByIdTodoList(ctx context.Context, in dtos.GetByIdTodoListInput) (dtos.GetByIdTodoListOutput, error)
}

func New(app app) (*echo.Echo, error) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	rest.New(e.Group("v1"), app)
	view.NewRouter(e.Group(""))

	return e, nil
}
