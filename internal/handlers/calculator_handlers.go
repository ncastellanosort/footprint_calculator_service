package handlers

import (
	"carbon_calculator/config"
	"encoding/json"
	"net/http"
	"sync"
)

type Message struct {
	Result float32 `json:"result"`
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request, respch chan float32, wg *sync.WaitGroup) {
	// manage CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	w.Header().Set("Content-type", "application/json")

	var answer config.Data
	json.NewDecoder(r.Body).Decode(&answer)
	json.NewEncoder(w).Encode(answer)
	/* 	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
	   		http.Error(w, "invalid request body", http.StatusBadRequest)
	   		return
	   		}
	   	   d	efer r.Body.Close()

	   	   	value := calc.Calculator(&answer, respch, wg)
	   	   	rounded_value := float32(math.Round(float64(value)*10) / 10)

	   	   	if err := json.NewEncoder(w).Encode(Message{Result: rounded_value}); err != nil {
	   	   		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	   	}
	*/
}
