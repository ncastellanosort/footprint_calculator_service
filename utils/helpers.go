package utils

import (
	"carbon_calculator/config"
	"fmt"
	"log"
	"sync"
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

func AnswersToArray(d map[string]int, k1 string, k2 string, k3 string, k4 string, convertArrayCh chan []float32, wg *sync.WaitGroup) {
	defer wg.Done()
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
	convertArrayCh <- res
}

func GetAnswers(answer *config.Data, convertArrayCh chan []float32, wg *sync.WaitGroup) (*config.Answers, error) {

	wg.Add(4)

	go AnswersToArray(answer.Energy, "applianceHours", "lightBulbs", "gasTanks", "hvacHours", convertArrayCh, wg)
	go AnswersToArray(answer.Waste, "trashBags", "foodWaste", "plasticBottles", "paperPackages", convertArrayCh, wg)
	go AnswersToArray(answer.Transport, "carKm", "publicKm", "domesticFlights", "internationalFlights", convertArrayCh, wg)
	go AnswersToArray(answer.Food, "redMeat", "whiteMeat", "dairy", "vegetarian", convertArrayCh, wg)

	wg.Wait()
	close(convertArrayCh)

	var results [][]float32

	for value := range convertArrayCh {
		results = append(results, value)
	}

	log.Println(results)

	if len(results) != 4 {
		return nil, fmt.Errorf("there are not 4 lists, are %d", len(results))
	}

	return &config.Answers{
		Transport: results[0],
		Energy:    results[1],
		Waste:     results[2],
		Food:      results[3],
	}, nil
}
