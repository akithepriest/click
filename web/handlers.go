package web

import (
	"net/http"

	"github.com/akithepriest/click/database"
	"github.com/labstack/echo/v4"
)

// Handler binds http handler functions.
// Handler struct accepts database connection pool.
type Handler struct {
	db *database.PostgresDB
}

func NewHandler(db *database.PostgresDB) *Handler {
	return &Handler{
		db: db,
	}
}

// Bind routes to their respective http handler.
func (h *Handler) DefineRoutes(server *echo.Echo) {
	server.GET("/", h.handleGET)
}

func (h *Handler) handleGET(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
