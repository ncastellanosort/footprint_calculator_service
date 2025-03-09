package routes

import (
	"carbon_calculator/internal/handlers"
	"net/http"
	"sync"
)

func SetupCalculatorRoutes(respch chan float32, wg *sync.WaitGroup) {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		handlers.CalculatorHandler(w, r, respch, wg)
	})
}
