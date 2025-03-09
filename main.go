package main

import (
	"carbon_calculator/config"
	"carbon_calculator/internal/calc"
)

var test_data config.Answers = config.Answers{
	Transport: []float32{130, 20, 1, 2},
	Food:      []float32{130, 20, 1, 2},
	Energy:    []float32{130, 20, 1, 2},
	Waste:     []float32{130, 20, 1, 2},
}

func main() {

	calc.Calculator(&test_data)

}
