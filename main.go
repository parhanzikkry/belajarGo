package main

import (
	"log"
	"net/http"

	"belajarGo/config"
	"belajarGo/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	cfg := config.GetConfig()

	// Init the route
	router := httprouter.New()
	handler.Init()
	handler.Route(router)

	// Listen and serve the server
	log.Fatal(http.ListenAndServe(":"+cfg.ServerConfig.Port, router))
}
