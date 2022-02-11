// Vamos criar uma loja, com vários produtos diferentes,
// onde as pessoas podem acessar e ver uma lista com todos
// meus produtos e alguns detalhes deles, como o nome, descrição, preço e quantidade.
package main

import (
	"net/http"

	"github.com/luciormoraes/web-app-100days/routes"
)

func main() {

	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
