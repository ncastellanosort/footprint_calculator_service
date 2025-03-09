package calc

import (
	"carbon_calculator/config"
	"sync"
)

func Calculator(data *config.Answers, respch chan float32, wg *sync.WaitGroup) float32 {
	transport_value := CalculateTransport(data, respch, wg)
	food_value := CalculateFood(data, respch, wg)
	waste_value := CalculateWaste(data, respch, wg)
	energy_value := CalculateEnergy(data, respch, wg)

	footprint := (transport_value + food_value + waste_value + energy_value)
	return footprint
}
