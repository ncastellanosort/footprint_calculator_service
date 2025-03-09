package main

import (
	"carbon_calculator/config"
	"carbon_calculator/internal/calc"
	"fmt"
)

var test_data config.Answers = config.Answers{
	Transport: []float32{130, 20, 1, 2},
	Food:      []float32{2, 3, 4, 0},
	Energy:    []float32{12, 8, 2, 0},
	Waste:     []float32{2, 1, 1, 3},
}

func main() {
	fmt.Println(fmt.Sprintf("Huella: %.1f CO2e/semana", calc.Calculator(&test_data)))
}
