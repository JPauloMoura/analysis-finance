package app

import (
	"net/http"

	"github.com/analysis-finance/app/handlers"
)

// StartServer sobe um servidor http
func StartServer() {
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/upload", handlers.HandleUpload)
}
