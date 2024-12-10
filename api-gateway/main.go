package main

import (
	"api-gateway/routes"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	// Register routes
	routes.RegisterProductRoutes(router)
	routes.RegisterOrderRoutes(router)

	log.Println("API Gateway running on port 8080")
	http.ListenAndServe(":8080", router)
}
