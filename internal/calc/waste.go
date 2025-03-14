package calc

import (
	"carbon_calculator/config"
	"carbon_calculator/utils"
	"sync"
)

func CalculateWaste(answers []float32, respch chan float32, wg *sync.WaitGroup) {
	defer wg.Done()
	value := utils.MultiplyAnswersAndEF(answers, config.EmisionFactors.WasteEmission)
	respch <- value
}
