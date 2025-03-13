package calc

import (
	"carbon_calculator/config"
	"carbon_calculator/utils"
	"sync"
)

func CalculateTransport(answers *config.Answers, respch chan float32, wg *sync.WaitGroup) {
	defer wg.Done()
	value := utils.MultiplyAnswersAndEF(answers.Transport, config.EmisionFactors.TransportEmission)
	respch <- value
}

/*
func CalculateT(data *config.Data, respch chan float32, wg *sync.WaitGroup) {
	defer wg.Done()
	value := utils.MultiplyAnswersAndEF(data.Transport, config.EmisionFactors.TransportEmission)
	respch <- value
}
*/
