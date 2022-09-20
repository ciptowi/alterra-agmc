package response

import (
	"github.com/labstack/echo/v4"
)

func ErrorResponse(c echo.Context, code int, message string, err error) error {
	res := new(RespError)
	res.Message = message
	res.Error = err
	return c.JSON(code, res)
}
