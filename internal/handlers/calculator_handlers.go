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
	Result float32 `json:"result"`
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request, respch chan float32, wg *sync.WaitGroup) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var answer config.Answers

	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	value := calc.Calculator(&answer, respch, wg)
	rounded_value := float32(math.Round(float64(value)*10) / 10)

	if err := json.NewEncoder(w).Encode(Message{Result: rounded_value}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
