package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"simple-todo-list/internal/adapters/api/rest"
	"simple-todo-list/internal/adapters/api/rest/app"
	"simple-todo-list/internal/view"
)

func New(app app.RestApp) (*echo.Echo, error) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	rest.New(e.Group("v1"), app)
	view.NewRouter(e.Group(""), app)

	return e, nil
}
