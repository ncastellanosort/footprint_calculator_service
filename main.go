package main

import (
	"carbon_calculator/internal/routes"
	"log"
	"net/http"
)

func main() {
	routes.SetupCalculatorRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
