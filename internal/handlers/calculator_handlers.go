package handlers

import (
	"carbon_calculator/config"
	"carbon_calculator/internal/calc"
	"encoding/json"
	"math"
	"net/http"
	"sync"
)

type Message struct {
	Result float32
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request, respch chan float32, wg *sync.WaitGroup) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-type", "application/json")

	var answer config.Answers

	json.NewDecoder(r.Body).Decode(&answer)

	value := calc.Calculator(&answer, respch, wg)
	rounded_value := float32(math.Round(float64(value)*10) / 10)

	json.NewEncoder(w).Encode(Message{Result: rounded_value})

}
