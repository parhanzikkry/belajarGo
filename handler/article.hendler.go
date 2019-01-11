package handler

import (
	"belajarGo/handler/response"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetAllArticle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := response.New()

	data, err := carticle.GetAllArticle()
	if err != nil {
		res.WriteError(w, r, http.StatusBadRequest, "Error when get article data")
		return
	}

	res.Data = data
	res.Status = "OK"
	res.WriteJson(w, r, http.StatusOK)

	return
}

func GetArticleById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := response.New()

	// get article id in url parameters
	articleId, _ := strconv.ParseInt(p.ByName("article_id"), 10, 64)
	if articleId <= 0 {
		res.WriteError(w, r, http.StatusUnprocessableEntity, "invalid product id")
		return
	}

	// get article data
	data, err := carticle.GetArticleById(articleId)
	if err != nil {
		res.WriteError(w, r, http.StatusBadRequest, "Error when get article data")
		return
	}

	res.Data = data
	res.WriteJson(w, r, http.StatusOK)

	return
}
