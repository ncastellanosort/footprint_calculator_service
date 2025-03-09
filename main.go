package main

import (
	"carbon_calculator/config"
	"carbon_calculator/internal/calc"
	"carbon_calculator/utils"
	"fmt"
)

var test_data config.Answers = config.Answers{
	Transport: []float32{130, 20, 1, 2},
	Food:      []float32{130, 20, 1, 2},
	Energy:    []float32{130, 20, 1, 2},
	Waste:     []float32{130, 20, 1, 2},
}

func Calculator(data *config.Answers) {
	// transport
	calc.CalculateTransport(&test_data)
	// food
	fmt.Println("food")
	// energy
	fmt.Println("energy")
	// waste
	fmt.Println("waste")

}

func main() {

	// 	Calculator(&test_data)
	// transport
	value := utils.MultiplyAnswersAndEF([]float32{2, 2, 3, 5}, config.EmisionFactors.TransportEmission)
	fmt.Println(value)

}
