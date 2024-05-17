package controller

import (
	inittemplate "boutique/templates"
	"net/http"
)

const Port = "localhost:8080"

var (
	query string
)

func NotFoundPageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	inittemplate.Temp.ExecuteTemplate(w, "404", nil)
}
