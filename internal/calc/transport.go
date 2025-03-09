package calc

import (
	"carbon_calculator/config"
	"carbon_calculator/utils"
	"fmt"
)

func CalculateTransport(answers *config.Answers) {
	fmt.Println("hello from CalculateTransport")
	utils.MultiplyAnswersAndEF(answers.Transport)
}
