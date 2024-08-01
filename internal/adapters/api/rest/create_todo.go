package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-todo-list/internal/dtos"
)

func createTodoList(app app) func(c echo.Context) error {
	return func(c echo.Context) error {
		var todoList dtos.CreateTodoListInput
		err := c.Bind(&todoList)
		if err != nil {
			return err
		}

		err = c.Validate(&todoList)
		if err != nil {
			return err
		}

		res, err := app.CreateTodoList(c.Request().Context(), todoList)
		if err != nil {
			return err
		}

		c.Response().Header().Add("Location", "/"+res.Id)

		return c.NoContent(http.StatusCreated)

	}
}
