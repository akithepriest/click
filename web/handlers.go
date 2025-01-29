package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler binds http handler functions.
// More fields will be added in this struct later.
type Handler struct {}

func NewHandler() *Handler {
	return &Handler{}
}

// Bind routes to their respective http handler.
func (h *Handler) DefineRoutes(server *echo.Echo) {
	server.GET("/", h.handleGET)
}

func (h *Handler) handleGET(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}