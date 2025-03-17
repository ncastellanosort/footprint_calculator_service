package handlers

import (
	"carbon_calculator/config"
	"carbon_calculator/internal/calc"
	"carbon_calculator/utils"
	"encoding/json"
	"math"
	"net/http"
	"sync"
)

type Message struct {
	Result float32 `json:"result"`
}

type DataMessage struct {
	Data   config.Data `json:"data"`
	Result float32     `json:"result"`
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request, calculateCh chan float32, wg *sync.WaitGroup, convertArrayCh chan config.ArrayData) {
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
	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
		http.Error(w, "invalid json payload", http.StatusBadRequest)
		return
	}

	answers, err := utils.GetAnswers(&answer, convertArrayCh, wg)
	if err != nil {
		http.Error(w, "failed getting answers", http.StatusInternalServerError)
		return
	}

	value, err := calc.Calculator(answers, calculateCh, wg)
	if err != nil {
		http.Error(w, "calculate error", http.StatusInternalServerError)
		return
	}
	rounded_value := float32(math.Round(float64(value)*10) / 10)

	if err := json.NewEncoder(w).Encode(DataMessage{Data: answer, Result: rounded_value}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}

}
