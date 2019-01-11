package routes

import (
	"github.com/julienschmidt/httprouter"
	Controller "belajarGo/controllers"
	Middleware "belajarGo/middlewares"
)

func Route(router *httprouter.Router) {
    router.GET("/", Middleware.CorsHandler(Controller.Index))
    router.GET("/hello/:name", Middleware.CorsHandler(Controller.Hello))
}