package router

//
//import (
//	"github.com/labstack/echo/v4"
//	"net/http"
//)
//
//type CustomValidator struct {
//	validator *vali
//}
//
//func (cv *CustomValidator) Validate(i interface{}) error {
//	if err := cv.validator.Struct(i); err != nil {
//		// Optionally, you could return the error to give each route more control over the status code
//		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//	}
//	return nil
//}
