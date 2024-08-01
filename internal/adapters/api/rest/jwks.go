package rest

import (
	"github.com/labstack/echo/v4"
	"simple-todo-list/internal/adapters/api/rest/app"
)

func getJwks(app app.RestApp) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.File("./pkg/auth/jwks.json")
	}
}
