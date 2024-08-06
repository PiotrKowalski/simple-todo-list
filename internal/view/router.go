package view

import "github.com/labstack/echo/v4"

func NewRouter(e *echo.Group, app app) *echo.Group {
	e.GET("", createIndexPageHandler())
	e.POST("", createUpdateTimeHandler())
	e.GET("/login", createLoginPageHandler(app))
	e.POST("/login", createLoginPageActionHandler(app))
	return e
}
