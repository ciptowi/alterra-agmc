package book

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create)
	g.PUT("/:id", h.UpdateByID)
	g.DELETE("/:id", h.DeleteByID)
	//g.GET("/:id", h.GetByID, middleware.Authentication)
}
