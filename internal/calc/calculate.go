package calc

import (
	"carbon_calculator/types"
	"sync"
)

func Calculator(data *types.Answers, respch chan float32, wg *sync.WaitGroup) (float32, error) {
	wg.Add(4)

	go CalculateTransport(data.Transport, respch, wg)
	go CalculateFood(data.Food, respch, wg)
	go CalculateWaste(data.Waste, respch, wg)
	go CalculateEnergy(data.Energy, respch, wg)

	wg.Wait()
	close(respch)

	var footprint float32

	for value := range respch {
		footprint += value
	}

	return footprint, nil
}
