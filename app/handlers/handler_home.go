package handlers

import (
	"html/template"
	"net/http"
)

var home = template.Must(template.ParseGlob("app/web/home.html"))

func HandleHome(w http.ResponseWriter, r *http.Request) {
	home.ExecuteTemplate(w, "Home", nil)
}
