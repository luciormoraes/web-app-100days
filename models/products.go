package models

import (
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
	selectAllProducts, err := db.Query("SELECT * FROM PRODUCTS")
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
