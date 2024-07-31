package router

import (
	"github.com/labstack/echo/v4"
	"log"
	"simple-todo-list/internal/dtos"
)

func getByIdTodoList(app app) func(c echo.Context) error {
	return func(c echo.Context) error {
		var todoList dtos.GetByIdTodoListInput
		err := c.Bind(&todoList)
		if err != nil {
			return err
		}
		log.Println(todoList)

		res, err := app.GetByIdTodoList(c.Request().Context(), todoList)
		if err != nil {
			return err
		}

		return c.JSON(200, res)

	}
}
