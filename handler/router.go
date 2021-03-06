package handler

import (
	"belajarGo/handler/middlewares"

	"github.com/julienschmidt/httprouter"
)

func Route(router *httprouter.Router) {
	router.GET("/article/:article_id", middlewares.SetHeader(GetArticleById))
	router.GET("/articel/all", middlewares.SetHeader(GetAllArticle))
}
