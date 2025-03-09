package calc

import (
	"carbon_calculator/config"
	"sync"
)

func Calculator(data *config.Answers, respch chan float32, wg *sync.WaitGroup) float32 {

	wg.Add(4)

	go CalculateTransport(data, respch, wg)
	go CalculateFood(data, respch, wg)
	go CalculateWaste(data, respch, wg)
	go CalculateEnergy(data, respch, wg)

	wg.Wait()
	close(respch)
	var footprint float32

	for value := range respch {
		footprint += value
	}
	return footprint
}
