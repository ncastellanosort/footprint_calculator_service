package calc

import (
	"carbon_calculator/config"
	"carbon_calculator/utils"
)

func CalculateWaste(answers *config.Answers) float32 {
	value := utils.MultiplyAnswersAndEF(answers.Waste, config.EmisionFactors.WasteEmission)
	return value
}
