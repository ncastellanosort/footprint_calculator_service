package internal

import (
	"carbon_calculator/types"
	"net/http"
	"sync"
)

func SetupCalculatorRoutes(wg *sync.WaitGroup) {

	http.HandleFunc("/carbon_calculator_service", func(w http.ResponseWriter, r *http.Request) {
		calculateCh := make(chan float32, 4)
		arrayCh := make(chan types.ArrayData, 4)

		CalculatorHandler(w, r, calculateCh, wg, arrayCh)
	})

}
