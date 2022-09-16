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

	e.POST("/books", controllers.SaveBook)
	e.GET("/books", controllers.GetBooks)
	e.GET("/books/:id", controllers.GetBookById)
	e.PUT("/books/:id", controllers.UpdateBookById)
	e.DELETE("/books/:id", controllers.DeleteBookById)

	return e
}
