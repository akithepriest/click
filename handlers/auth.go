package handlers

import (
	"errors"
	"net/http"
	"os"

	"github.com/akithepriest/click/services"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	collection *mongo.Collection
	service *services.GoogleOAuthService
}

func NewAuthHandler(db *mongo.Database) (*AuthHandler, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackUrl := os.Getenv("GOOGLE_CLIENT_CALLBACK_URL")

	if clientID == "" || clientSecret == "" || callbackUrl == "" {
		return nil, errors.New("credentials for google are not defined in the environment")
	}

	collection := db.Collection("users")
	service := services.NewGoogleOAuthService(clientID, clientSecret, callbackUrl)
	return &AuthHandler{
		collection: collection,
		service: service,
	}, nil
}

func (h *AuthHandler) DefineRoutes(e *echo.Group) {
	e.GET("/google/login", h.handleGETLogin)
	e.GET("/google/callback", h.handleGETCallback)
}

func (h *AuthHandler) handleGETLogin(c echo.Context) error {
	u := h.service.GetLoginUrl(c.Response().Writer)
	return c.Redirect(http.StatusTemporaryRedirect, u)
}

func (h *AuthHandler) handleGETCallback(c echo.Context) error {
	code := c.FormValue("code")
	return c.String(http.StatusOK, "code is " + code)
}