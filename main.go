// Vamos criar uma loja, com vários produtos diferentes,
// onde as pessoas podem acessar e ver uma lista com todos
// meus produtos e alguns detalhes deles, como o nome, descrição, preço e quantidade.
package main

import (
	"html/template"
	"net/http"

	"github.com/luciormoraes/web-app-100days/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
