package utils

import (
	"carbon_calculator/config"
	"log"
)

func MultiplyAnswersAndEF(answers []float32, ef []float32) float32 {
	var res float32
	if len(answers) != len(ef) {
		log.Fatal("arrays must be the same length")
		return 0
	}

	for index, answer := range answers {
		res += answer * ef[index]
	}

	return res
}

func AnswersToArray(d map[string]int, k1 string, k2 string, k3 string, k4 string) []float32 {
	res := make([]float32, len(d))
	for key, value := range d {
		if key == k1 {
			res[0] = float32(value)
		}

		if key == k2 {
			res[1] = float32(value)
		}

		if key == k3 {
			res[2] = float32(value)
		}

		if key == k4 {
			res[3] = float32(value)
		}
	}
	return res
}
func GetAnswers(answer *config.Data) *config.Answers {

	energy := AnswersToArray(answer.Energy, "applianceHours", "lightBulbs", "gasTanks", "hvacHours")
	waste := AnswersToArray(answer.Waste, "trashBagsa", "foodWaste", "plasticBottles", "paperPackages")
	transport := AnswersToArray(answer.Transport, "carKm", "publicKm", "domesticFlights", "internationalFlights")
	food := AnswersToArray(answer.Food, "redMeat", "whiteMeat", "dairy", "vegetarian")

	return &config.Answers{
		Transport: transport,
		Energy:    energy,
		Waste:     waste,
		Food:      food,
	}
}
