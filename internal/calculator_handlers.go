package internal

import (
	"carbon_calculator/internal/auth"
	"carbon_calculator/internal/calc"
	"carbon_calculator/types"
	"encoding/json"
	"math"
	"net/http"
	"sync"
)

func CalculatorHandler(w http.ResponseWriter, r *http.Request, calculateCh chan float32, wg *sync.WaitGroup, convertArrayCh chan types.ArrayData) {
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

	token := r.Header.Get("Authorization")

	valid := auth.ValidateToken(token)

	var answer types.Data

	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
		http.Error(w, "invalid json payload", http.StatusBadRequest)
		return
	}

	// user logged
	if valid {
		processUserFlow(true, answer, token, w, calculateCh, wg, convertArrayCh)
	} else {
		// not logged
		processUserFlow(false, answer, token, w, calculateCh, wg, convertArrayCh)
	}

}

func processUserFlow(is_logged bool, answer types.Data, token string, w http.ResponseWriter, calculateCh chan float32, wg *sync.WaitGroup, convertArrayCh chan types.ArrayData) {
	value, err := calc.GetAnswers(is_logged, &answer, convertArrayCh, wg, calculateCh)
	if err != nil {
		http.Error(w, "failed getting answers", http.StatusInternalServerError)
		return
	}

	rounded := float32(math.Round(float64(value)*10) / 10)
	response := types.DataResponse{Data: answer, Result: rounded}

	/*
	res := utils.PostData(response, token)
	if res == nil {
		http.Error(w, "failed posting data", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	*/

	/*
	final := types.Message{
		//Status: res.StatusCode,
		Status: http.StatusOK,
		Info:   "data sent",
	}
	*/

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed encoding response", http.StatusInternalServerError)
	}
}
