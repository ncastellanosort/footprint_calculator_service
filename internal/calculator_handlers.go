package internal

import (
	"carbon_calculator/internal/auth"
	"carbon_calculator/internal/calc"
	"carbon_calculator/types"
	"carbon_calculator/utils"
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

	valid := auth.Validate_token(token)

	var answer types.Data

	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
		http.Error(w, "invalid json payload", http.StatusBadRequest)
		return
	}

	// user logged
	if valid {

		logged_answers, err := utils.GetAnswers(true, &answer, convertArrayCh, wg)
		if err != nil {
			http.Error(w, "failed getting answers", http.StatusInternalServerError)
			return
		}

		logged_value, err := calc.Calculator(logged_answers, calculateCh, wg)

		if err != nil {
			http.Error(w, "calculate error", http.StatusInternalServerError)
			return
		}

		rounded_value := float32(math.Round(float64(logged_value)*10) / 10)

		w.WriteHeader(http.StatusOK)

		res := utils.PostData(types.DataResponse{Data: answer, Result: rounded_value}, token)

		defer res.Body.Close()

		message := types.Message{Status:res.StatusCode, Info: "data sent"}

		if err := json.NewEncoder(w).Encode(message); err != nil {
			http.Error(w, "failed sending data", http.StatusInternalServerError)
		}

	} else {

		// user not logged
		not_logged_answers, err := utils.GetAnswers(false, &answer, convertArrayCh, wg)

		if err != nil {
			http.Error(w, "failed getting answers", http.StatusInternalServerError)
			return
		}

		not_logged_value, err := calc.Calculator(not_logged_answers, calculateCh, wg)

		if err != nil {
			http.Error(w, "calculate error", http.StatusInternalServerError)
			return
		}

		rounded_value := float32(math.Round(float64(not_logged_value)*10) / 10)

		w.WriteHeader(http.StatusOK)

		res := utils.PostData(types.DataResponse{Data: answer, Result: rounded_value}, token)

		defer res.Body.Close()

		message := types.Message{Status:res.StatusCode, Info: "data sent"}

		if err := json.NewEncoder(w).Encode(message); err != nil {
			http.Error(w, "failed sending data", http.StatusInternalServerError)
		}
	}

}
