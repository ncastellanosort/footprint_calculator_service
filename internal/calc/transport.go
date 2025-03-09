package calc

import (
	"carbon_calculator/config"
	"carbon_calculator/utils"
)

func CalculateTransport(answers *config.Answers) float32 {
	value := utils.MultiplyAnswersAndEF(answers.Transport, config.EmisionFactors.TransportEmission)
	return value
}
