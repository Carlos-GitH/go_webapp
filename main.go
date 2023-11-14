package main

import (
	"alura/webapp/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// db := conectaComBancoDeDados()
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
