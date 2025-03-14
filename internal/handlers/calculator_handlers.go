package handlers

import (
	"carbon_calculator/config"
	"carbon_calculator/utils"
	"encoding/json"
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

	energy := utils.AnswersToArray(answer.Energy)
	waste := utils.AnswersToArray(answer.Waste)
	transport := utils.AnswersToArray(answer.Transport)
	food := utils.AnswersToArray(answer.Food)

	answers := config.Answers{
		Transport: transport,
		Energy:    energy,
		Waste:     waste,
		Food:      food,
	}

	json.NewEncoder(w).Encode(answers)
	defer r.Body.Close()
	/*
			value := calc.Calculator(&answers, respch, wg)
		rounded_value := float32(math.Round(float64(value)*10) / 10)

			json.NewEncoder(w).Encode(DataMessage{Data: answer, Result: rounded_value})
	*/
}
