package main

import (
	"GoLangAppWeb/routes"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
