package calc

import (
	"carbon_calculator/config"
	"carbon_calculator/utils"
)

func CalculateEnergy(answers *config.Answers) float32 {
	value := utils.MultiplyAnswersAndEF(answers.Energy, config.EmisionFactors.EnergyEmission)
	return value
}
