package user

//
//import (
//	"github.com/labstack/echo/v4"
//	"net/http"
//	"simple-todo-list/internal/adapters/api/rest"
//	"simple-todo-list/internal/dtos/user"
//)
//
//func login(app rest.RestApp) func(c echo.Context) error {
//	return func(c echo.Context) error {
//		var loginInput user.LoginInput
//		err := c.Bind(&loginInput)
//		if err != nil {
//			return err
//		}
//
//		err = c.Validate(&loginInput)
//		if err != nil {
//			return err
//		}
//
//		res, err := app.CreateTodoList(c.Request().Context(), loginInput)
//		if err != nil {
//			return err
//		}
//
//		c.Response().Header().Add("Location", "/"+res.Id)
//
//		return c.NoContent(http.StatusCreated)
//
//	}
//}
