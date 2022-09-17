package restrict

import (
	"auth/controllers"
	"auth/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(controllers.JWT_SECRET),
})

func Authorization(c echo.Context) error {

	res := new(helpers.Response)
	if err := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(controllers.JWT_SECRET),
	}); err != nil {
		res.Message = "Can't accesss token"
		return c.JSON(http.StatusUnauthorized, res)
	}
	return nil
}
