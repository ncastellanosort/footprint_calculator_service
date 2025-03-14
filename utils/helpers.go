package utils

import (
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
	var res []float32
	for _, value := range d {
		res = append(res, float32(value))
	}
	return res
}
