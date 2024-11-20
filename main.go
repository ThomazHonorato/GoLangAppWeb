package main

import (
	"GoLangAppWeb/models"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()

	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}
