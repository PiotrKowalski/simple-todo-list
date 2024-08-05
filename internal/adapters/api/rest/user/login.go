package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-todo-list/internal/adapters/api/rest/app"
	"simple-todo-list/internal/dtos"
)

func login(app app.RestApp) func(c echo.Context) error {
	return func(c echo.Context) error {
		var loginInput dtos.LoginInput
		err := c.Bind(&loginInput)
		if err != nil {
			return err
		}

		err = c.Validate(&loginInput)
		if err != nil {
			return err
		}

		res, err := app.Login(c.Request().Context(), loginInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		//utils.AddLocationHeaderToResponse(c, res.Id)

		return c.JSON(http.StatusOK, res)

	}
}
