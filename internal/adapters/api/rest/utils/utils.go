package utils

import (
	"github.com/labstack/echo/v4"
	"simple-todo-list/pkg/config"
)

var (
	PORT, _     = config.ReadEnvString("APP_PORT")
	HOSTNAME, _ = config.ReadEnvString("HOSTNAME")
)

func getAbsolutePath() string {
	return HOSTNAME + ":" + PORT
}

func AddLocationHeaderToResponse(c echo.Context, value string) {
	c.Response().Header().Add("Location", getAbsolutePath()+"/"+value)
}
