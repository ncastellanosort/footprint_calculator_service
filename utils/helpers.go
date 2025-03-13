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

func AnswersToArray(data config.Data) []float32 {
	var res []float32
	// iterar los valores de cada uno con el key, value
	// creo que debe estar dentro de config y debe estar attach a la struct
	return res
}
