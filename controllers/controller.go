package controllers

import (
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"
)

type dataResponse struct {
	Status		bool	`json:"status"`
	Message		string	`json:"message"`
	Params		string	`json:"parameter"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	DataResponse := dataResponse {
		Status	: true,
		Message	: "Welcome!",
		Params	: "tidak ada",
	}
	data,_ := json.Marshal(DataResponse)
    w.Write(data)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	DataResponse := dataResponse {
		Status	: true,
		Message	: "Hello, " + ps.ByName("name"),
		Params	: ps.ByName("name"),
	}
	data,_ := json.Marshal(DataResponse)
    w.Write(data)
}