package http

import (
	"github.com/labstack/echo/v4"
	"library-sevice/internal/app/book"
	"library-sevice/internal/app/user"
	"library-sevice/internal/factory"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	user.NewHandler(f).Route(e.Group("/users"))
	book.NewHandler(f).Route(e.Group("/books"))
}
