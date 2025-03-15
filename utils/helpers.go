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

func AnswersToArray(d map[string]int) []float32 {
	res := make([]float32, len(d))
	n := 0
	for _, value := range d {
		res[n] = float32(value)
		n++
	}
	return res
}

func GetAnswers(answer *config.Data) *config.Answers {

	energy := AnswersToArray(answer.Energy)
	waste := AnswersToArray(answer.Waste)
	transport := AnswersToArray(answer.Transport)
	food := AnswersToArray(answer.Food)

	return &config.Answers{
		Transport: transport,
		Energy:    energy,
		Waste:     waste,
		Food:      food,
	}
}
