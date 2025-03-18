package calc

import (
	"carbon_calculator/types"
	"carbon_calculator/utils"
	"sync"
)

func CalculateWaste(answers []float32, respch chan float32, wg *sync.WaitGroup) {
	defer wg.Done()
	value := utils.MultiplyAnswersAndEF(answers, types.EmisionFactors.WasteEmission)
	respch <- value
}
