package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
	"integration-test/helpers"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(helpers.JWT_SECRET),
})
