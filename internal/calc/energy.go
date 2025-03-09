package calc

import (
	"carbon_calculator/config"
	"carbon_calculator/utils"
	"sync"
)

func CalculateEnergy(answers *config.Answers, respch chan float32, wg *sync.WaitGroup) float32 {
	value := utils.MultiplyAnswersAndEF(answers.Energy, config.EmisionFactors.EnergyEmission)
	return value
}
