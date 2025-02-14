package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LandingHandler struct{}

func (h LandingHandler) DefineRoutes(e *echo.Group) {
	e.GET("/", h.handleGET)
}

func (h LandingHandler) handleGET(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}