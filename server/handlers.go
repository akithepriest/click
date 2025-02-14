package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler binds http handler functions.
// Handler struct accepts database connection pool.
type Handler struct {
	db *mongo.Database
}

func NewHandler(db *mongo.Database) *Handler {
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