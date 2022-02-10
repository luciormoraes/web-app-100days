// Vamos criar uma loja, com vários produtos diferentes,
// onde as pessoas podem acessar e ver uma lista com todos
// meus produtos e alguns detalhes deles, como o nome, descrição, preço e quantidade.
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	// conn:="user dbname password host sslmode"
	conn := "user=root dbname=test_db password=root host=0.0.0.0 sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
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
	temp.ExecuteTemplate(w, "Index", products)
	fmt.Println(products)
	defer db.Close()
}
