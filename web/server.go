package web

import (
	"html/template"
	"os"

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

func (w *WebServer) BindHandlers() {
	handler := NewHandler()
	handler.DefineRoutes(w.server)
}

func (w *WebServer) Start() {
	listenAddr := os.Getenv("LISTEN_ADDRESS") 
	if listenAddr == "" {
		w.server.Logger.Fatal("LISTEN_ADDRESS is not defined in the environment")
	}
	w.server.Logger.Info(w.server.Start(listenAddr))
}