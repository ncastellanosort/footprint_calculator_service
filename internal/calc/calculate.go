package calc

import (
	"carbon_calculator/config"
)

func Calculator(data *config.Answers) float32 {
	transport_value := CalculateTransport(data)
	food_value := CalculateFood(data)
	waste_value := CalculateWaste(data)
	energy_value := CalculateEnergy(data)

	footprint := (transport_value + food_value + waste_value + energy_value)
	return footprint
}
