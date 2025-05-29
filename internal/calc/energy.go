package calc

import (
	"carbon_calculator/types"
	"sync"
)

func CalculateEnergy(answers []float32, respch chan float32, wg *sync.WaitGroup) {
	defer wg.Done()
	value := MultiplyAnswersAndEF(answers, types.EmisionFactors.EnergyEmission)
	respch <- value
}
