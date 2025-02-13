package main

import (
	"flag"
	"log"

	"github.com/akithepriest/click/server"
	"github.com/joho/godotenv"
)

var isDev = flag.Bool("dev", false, "Whether to start application in development mode")

func main() {
	flag.Parse()
	if *isDev {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatalln("Failed to load environmental variables from .env.local because ", err)
		}
	}
	
	server := server.NewWebServer()
	server.BindHandlers()
	server.Start()
}