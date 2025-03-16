package routes

import (
	"carbon_calculator/internal/handlers"
	"net/http"
	"sync"
)

func SetupCalculatorRoutes(wg *sync.WaitGroup) {

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		calculateCh := make(chan float32, 4)
		convertArrayCh := make(chan []float32, 4)

		handlers.CalculatorHandler(w, r, calculateCh, wg, convertArrayCh)
	})

}
