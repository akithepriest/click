package main

import "github.com/akithepriest/click/web"

func main() {
	server := web.NewWebServer()
	server.BindHandlers()
	server.Start()
}