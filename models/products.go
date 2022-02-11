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
	selectAllProducts, err := db.Query("select * from products")
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
