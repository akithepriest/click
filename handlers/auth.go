package handlers

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/akithepriest/click/database"
	"github.com/akithepriest/click/middlewares"
	"github.com/akithepriest/click/services"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	collection *mongo.Collection
	oauthService *services.GoogleOAuthService
	userService *services.UserService
}

func NewAuthHandler(db *mongo.Database) (*AuthHandler, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackUrl := os.Getenv("GOOGLE_CLIENT_CALLBACK_URL")

	if clientID == "" || clientSecret == "" || callbackUrl == "" {
		return nil, errors.New("credentials for google are not defined in the environment")
	}

	collection := db.Collection("users")
	oauthService := services.NewGoogleOAuthService(clientID, clientSecret, callbackUrl)
	userService := services.NewUserService(collection)

	return &AuthHandler{
		collection: collection,
		oauthService: oauthService,
		userService: userService,
	}, nil
}

func (h *AuthHandler) DefineRoutes(e *echo.Group) {
	e.GET("/google/login", h.handleGETLogin)
	e.GET("/google/callback", h.handleGETCallback)
	e.GET("/protected", middlewares.ProtectedMiddleware(h.handleProtected))
}

func (h *AuthHandler) handleGETLogin(c echo.Context) error {
	u := h.oauthService.GetLoginUrl()
	return c.Redirect(http.StatusTemporaryRedirect, u)
}

// User is redirected to this route after google auth
func (h *AuthHandler) handleGETCallback(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	code := c.FormValue("code")
	if code == "" {
		return c.String(http.StatusBadRequest, "code is missing")
	}
	userinfo, err := h.oauthService.GetUserData(ctx, code)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	user, err := h.userService.InsertUser(ctx, userinfo.Name, userinfo.Email)
	if err != nil {
		if errors.Is(err, database.ErrorAlreadyExists) {
			return c.String(http.StatusOK, "user is already registered")
		}
	}
	
	tokenString, err := services.CreateJWT(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	services.SetJWTCookie(c.Response().Writer, tokenString)
	return c.String(http.StatusOK, "email: " + user.Email + ", name:" + user.Name + " id:" + user.ID.String())
}

func (h *AuthHandler) handleProtected(c echo.Context) error {
	return c.String(http.StatusOK, "You have permission to view this path")
}