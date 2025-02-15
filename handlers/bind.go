package handlers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler interface {
	DefineRoutes(e *echo.Group)
}

func BindHandlers(e *echo.Echo, db *mongo.Database) {
	indexGroup := e.Group("")
	LandingHandler{}.DefineRoutes(indexGroup)

	authGroup := e.Group("/auth")
	if authHandler, err := NewAuthHandler(db); err != nil {
		e.Logger.Error("error in auth handler: ", err)
	} else {
		authHandler.DefineRoutes(authGroup)
	}

}