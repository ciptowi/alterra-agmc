package restrict

import (
	"integration-test/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(helpers.JWT_SECRET),
})

func Authorization(c echo.Context) error {

	res := new(helpers.Response)
	if err := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(helpers.JWT_SECRET),
	}); err != nil {
		res.Message = "Can't accesss token"
		return c.JSON(http.StatusUnauthorized, res)
	}
	return nil
}
