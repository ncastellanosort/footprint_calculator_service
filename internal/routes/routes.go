package routes

import (
	"carbon_calculator/internal/handlers"
	"net/http"
)

func SetupCalculatorRoutes() {
	http.HandleFunc("/calculate", handlers.CalculatorHandler)
}
