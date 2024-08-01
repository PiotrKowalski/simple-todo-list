package rest

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"simple-todo-list/internal/adapters/api/rest/app"
	api_todolist "simple-todo-list/internal/adapters/api/rest/todolist"
	api_user "simple-todo-list/internal/adapters/api/rest/user"
	"simple-todo-list/pkg/auth"
)

func New(e *echo.Group, app app.RestApp) {
	e.GET("/jwks", getJwks(app))
	api_user.NewRouter(e.Group("/user"), app)

	restricted := e.Group("")
	restricted.Use(echojwt.WithConfig(echojwt.Config{KeyFunc: auth.FetchJwks}))
	api_todolist.NewRouter(restricted.Group("/todolist"), app)
}
