package routers

import (
	"integration-test/controllers"
	"integration-test/restrict"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(restrict.CostumLogger))

	e.POST("/users", controllers.SaveUser)
	e.GET("/users", controllers.GetUsers, restrict.IsAuthenticated)
	e.GET("/users/:id", controllers.GetUserById, restrict.IsAuthenticated)
	e.PUT("/users/:id", controllers.UpdateUserById, restrict.IsAuthenticated)
	e.DELETE("/users/:id", controllers.DeleteUserById, restrict.IsAuthenticated)

	e.POST("/books", controllers.SaveBook, restrict.IsAuthenticated)
	e.GET("/books", controllers.GetBooks)
	e.GET("/books/:id", controllers.GetBookById)
	e.PUT("/books/:id", controllers.UpdateBookById, restrict.IsAuthenticated)
	e.DELETE("/books/:id", controllers.DeleteBookById, restrict.IsAuthenticated)

	e.POST("/login", controllers.LoginByEmailAndPassword)

	return e
}
