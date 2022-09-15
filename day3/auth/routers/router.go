package routers

import (
	"auth/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.POST("/users", controllers.SaveUser)
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUserById)
	e.PUT("/users/:id", controllers.UpdateUserById)
	e.DELETE("/users/:id", controllers.DeleteUserById)
	return e
}
