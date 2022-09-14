package routes

import (
	"github.com/labstack/echo/v4"
	"static/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", controllers.Index)
	e.POST("/books", controllers.SaveBook)
	e.GET("/books", controllers.GetBooks)
	e.GET("/books/:id", controllers.GetBookById)
	e.PUT("/books/:id", controllers.UpdateBookById)
	e.DELETE("/books/:id", controllers.DeleteBookById)

	return e
}
