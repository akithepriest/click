package web

import "github.com/labstack/echo/v4"

type WebServer struct {
	server *echo.Echo
}

func NewWebServer() *WebServer {
	server := echo.New()
	return &WebServer{
		server: server,
	}
}

func (w *WebServer) BindHandlers() {
	handler := NewHandler()
	handler.DefineRoutes(w.server)
}

func (w *WebServer) Start() {
	w.server.Logger.Info(w.server.Start(":8080"))
}