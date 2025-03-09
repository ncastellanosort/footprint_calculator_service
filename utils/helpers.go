package utils

func MultiplyAnswersAndEF(answers []float32, ef []float32) float32 {
	var res float32
	for index, answer := range answers {
		res += answer * ef[index]
	}

	return res
}
