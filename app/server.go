package app

import (
	"net/http"

	"github.com/analysis-finance/app/handlers"
)

func StartServer() {
	http.HandleFunc("/", handlers.HandleHome)
}
