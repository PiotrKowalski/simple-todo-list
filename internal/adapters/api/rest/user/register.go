package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-todo-list/internal/adapters/api/rest/app"
	"simple-todo-list/internal/adapters/api/rest/utils"
	"simple-todo-list/internal/dtos"
)

func Register(app app.RestApp) func(c echo.Context) error {
	return func(c echo.Context) error {
		var registerInput dtos.RegisterInput
		err := c.Bind(&registerInput)
		if err != nil {
			return err
		}

		err = c.Validate(&registerInput)
		if err != nil {
			return err
		}

		res, err := app.Register(c.Request().Context(), registerInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		utils.AddLocationHeaderToResponse(c, res.Id)

		return c.JSON(http.StatusOK, res)

	}
}
