package models

import (
	"fmt"

	"github.com/luciormoraes/web-app-100days/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SearchAllProducts() []Product {
	db := db.ConnectDB()
	selectAllProducts, err := db.Query("SELECT * FROM PRODUCTS ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	// products := []Product{
	// 	{Name: "T-shirt", Description: "Vader", Price: 39, Quantity: 10},
	// 	{"Ride Boots", "HD", 300, 12},
	// 	{"Gimbal", "SpiceyX", 109.54, 10},
	// 	{"Stand Desk", "Black Matrix", 359, 5},
	// }
	defer db.Close()
	return products
}

func SaveNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDB()
	insertData, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(idProduct string) {
	db := db.ConnectDB()
	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteProduct.Exec(idProduct)
	defer db.Close()
}

func EditProduct(idProduct string) Product {
	db := db.ConnectDB()
	selectProductById, err := db.Query("SELECT * FROM products WHERE id=$1", idProduct)
	if err != nil {
		panic(err.Error())
	}
	productToUpdate := Product{}

	for selectProductById.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProductById.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity

	}
	fmt.Println(productToUpdate)
	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.ConnectDB()
	updateProduct, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
