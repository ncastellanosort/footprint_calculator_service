package calc

import (
	"carbon_calculator/types"
	"sync"
)

func CalculateWaste(answers []float32, respch chan float32, wg *sync.WaitGroup) {
	defer wg.Done()
	value := MultiplyAnswersAndEF(answers, types.EmisionFactors.WasteEmission)
	respch <- value
}
