package main

import (
	"log"
	"net/http"

	"github.com/analysis-finance/app"
)

func main() {
	app.StartServer()
	log.Println("Servidor rodando na porta 3002...")
	http.ListenAndServe(":3002", nil)
}
