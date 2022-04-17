package handlers

import (
	"html/template"
	"net/http"
)

var home = template.Must(template.ParseGlob("app/views/home.html"))

// HandleHome retorna a página inicial de upload de arquivos
func HandleHome(w http.ResponseWriter, r *http.Request) {
	home.ExecuteTemplate(w, "Home", nil)
}
