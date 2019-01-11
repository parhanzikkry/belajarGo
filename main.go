package main

import (
	"belajarGo/routes"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Get PORT from env of device or server
	// port := os.Getenv("PORT")

	// Init the route
	router := httprouter.New()
	routes.Route(router)

	// Listen and serve the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
