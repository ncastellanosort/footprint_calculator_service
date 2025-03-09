package calc

import (
	"carbon_calculator/config"
	"carbon_calculator/utils"
)

func CalculateFood(answers *config.Answers) float32 {
	value := utils.MultiplyAnswersAndEF(answers.Food, config.EmisionFactors.FoodEmission)
	return value
}
