package todolist

import (
	"github.com/labstack/echo/v4"
	"simple-todo-list/internal/adapters/api/rest/app"
)

func NewRouter(e *echo.Group, app app.RestApp) {
	e.POST("", createTodoList(app))
	e.GET("/:id", getByIdTodoList(app))

}
