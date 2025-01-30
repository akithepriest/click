package web

import (
	"context"
	"errors"
	"html/template"
	"os"
	"time"

	"github.com/akithepriest/click/database"
	"github.com/labstack/echo/v4"
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
	
	return &WebServer{
		server: server,
	}
}

func (w *WebServer) initDB() (*database.PostgresDB, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 30)
	defer cancel()

	connString := os.Getenv("PG_DATABASE_URL")
	if connString == "" {
		return nil, errors.New("failed to connect to postgresql database, could not find PG_DATABASE_URL in environment")
	}
	db, err := database.NewPgDB(ctx, connString)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("db: PostgresDB is null")
	}
	return db, nil
}

func (w *WebServer) BindHandlers() {
	db, err := w.initDB()
	if err != nil {
		w.server.Logger.Fatal(err)
		return 
	}
	w.server.Logger.Info("Connection to database has been established.")

	handler := NewHandler(db)
	handler.DefineRoutes(w.server)
	w.server.Logger.Info("Handlers have been registered.")
}

func (w *WebServer) Start() {
	listenAddr := os.Getenv("LISTEN_ADDRESS") 
	if listenAddr == "" {
		w.server.Logger.Fatal("LISTEN_ADDRESS is not defined in the environment")
	}
	w.server.Logger.Info(w.server.Start(listenAddr))
}