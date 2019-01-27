package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type (
	Response struct {
		Status         string      `json:"status, omitempty"`
		Data           interface{} `json: "data, omitempty"`
		Errors         interface{} `json: "errors, omitempty"`
		ProcessTime    string      `json: "server_process_time"`
		Server         string      `json: "server"`
		Token          string      `json: "token, omitempty"`
		ServerTimeUnix int64       `json: "server_time_unix"`

		origin    string
		startTime time.Time
	}

	Links struct {
		Self string `json:"self,omitempty"`
		Next string `json:"next,omitempty"`
		Prev string `json:"prev,omitempty"`
	}
)

func New(origin ...string) Response {
	r := Response{
		startTime: time.Now(),
		Server:    "Localhost",
	}

	return r
}

func (res *Response) WriteJson(w http.ResponseWriter, r *http.Request, status int) {

	// make sure start time not nil
	if res.startTime.IsZero() {
		res.startTime = time.Now()
	}

	// calculate process time
	if res.ProcessTime == "" {
		res.ProcessTime = fmt.Sprintf("%.2f ms", time.Since(res.startTime).Seconds()*1000)
	}

	// server time
	res.ServerTimeUnix = time.Now().Unix()

	jsonByte, err := json.Marshal(res)
	if err != nil {
		log.Printf("[Response] WriteJson fail in json.Marshal")
	}

	w.Write(jsonByte)
}

func (res *Response) WriteError(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	res.Errors = message
	res.WriteJson(w, r, status)
	return
}
