package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{
			Nome:       "Camiseta",
			Descricao:  "Bem Bonita",
			Preco:      39,
			Quantidade: 5,
		},
		{
			"Tenis",
			"Confortavel",
			89,
			3,
		},
		{
			"Fone",
			"Muito Bom",
			59,
			2,
		},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}
