package main

import (
	"carbon_calculator/internal/routes"
	"fmt"
	"log"
	"net/http"
)

const PORT int = 8080

func main() {
	routes.SetupCalculatorRoutes()
	fmt.Println(fmt.Sprintf("server running in http://localhost:%d", PORT))
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
