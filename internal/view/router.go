package view

import "github.com/labstack/echo/v4"

func NewRouter(e *echo.Group) *echo.Group {
	e.GET("", createIndexPageHandler())
	e.POST("", createUpdateTimeHandler())
	e.GET("/login", createLoginPageHandler())
	return e
}
