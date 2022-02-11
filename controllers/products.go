package controllers

import (
	"net/http"
	"text/template"

	"github.com/luciormoraes/web-app-100days/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
