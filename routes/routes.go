package routes

import (
	"net/http"

	"github.com/luciormoraes/web-app-100days/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
