package internal

import (
	"bytes"
	"carbon_calculator/types"
	"carbon_calculator/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func Test_calculator_handler(t *testing.T) {

	utils.Connect()

	ch1 := make(chan float32, 4)
	ch2 := make(chan types.ArrayData, 4)
	wg := &sync.WaitGroup{}

	testData := types.Data{
		Date: "2025-03-19",
		Energy: map[string]int{
			"applianceHours": 12,
			"lightBulbs":     4,
			"gasTanks":       15,
			"hvacHours":      20,
		},
		Food: map[string]int{
			"redMeat":    25,
			"whiteMeat":  20,
			"dairy":      20,
			"vegetarian": 2,
		},
		Transport: map[string]int{
			"carKm":                230,
			"publicKm":             150,
			"domesticFlights":      5,
			"internationalFlights": 2,
		},
		Waste: map[string]int{
			"trashBags":      5,
			"foodWaste":      5,
			"plasticBottles": 6,
			"paperPackages":  5,
		},
	}

	jsn_body, _ := json.Marshal(testData)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "/carbon_calculator_service", bytes.NewBuffer(jsn_body))
	r.Header.Set("Content-type", "application/json")

	CalculatorHandler(w, r, ch1, wg, ch2)

	defer r.Body.Close()

	// body format
	// method
	// expected response value

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("expected %d got %d", http.StatusOK, w.Result().StatusCode)
	}

}
