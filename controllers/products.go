package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/luciormoraes/web-app-100days/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// From form
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")
		// Convert from string
		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro while try to convert price:", err)
		}
		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro while try to convert quantity:", err)
		}

		models.SaveNewProduct(name, description, priceConverted, quantityConverted)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	productToUpdate := models.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Update", productToUpdate)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro while try to convert ID to int:", err)
		}
		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro while try to convert PRICE to float:", err)
		}
		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro while try to convert QUANTITY to int:", err)
		}

		models.UpdateProduct(idConverted, name, description, priceConverted, quantityConverted)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
