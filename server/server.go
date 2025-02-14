package server

import (
	"context"
	"errors"
	"html/template"
	"os"
	"time"

	"github.com/akithepriest/click/database"
	"github.com/akithepriest/click/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type WebServer struct {
	server *echo.Echo
}

var t = TemplatesRenderer{
	templates: template.Must(template.ParseGlob("public/views/*.html")),
}

func NewWebServer() *WebServer {
	server := echo.New()
	server.Renderer = &t
	server.Logger.SetLevel(log.INFO)
	
	return &WebServer{
		server: server,
	}
}

func (w *WebServer) initDB() (*mongo.Database, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	client, err := database.NewMongoClient(ctx)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, errors.New("db: mongo.Client is null")
	}

	database := client.Database("master")
	return database, nil
}

func (w *WebServer) BindHandlers() {
	db, err := w.initDB()
	if err != nil {
		w.server.Logger.Fatal(err)
		return 
	}
	w.server.Logger.Info("Connection to database has been established.")

	handlers.BindHandlers(w.server, db)
	w.server.Logger.Info("Handlers have been registered.")
}

func (w *WebServer) Start() {
	listenAddr := os.Getenv("LISTEN_ADDRESS") 
	if listenAddr == "" {
		w.server.Logger.Fatal("LISTEN_ADDRESS is not defined in the environment")
	}
	w.server.Logger.Info(w.server.Start(listenAddr))
}