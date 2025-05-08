package utils

import (
	"carbon_calculator/types"
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

func SumAnswers(answers []float32) float32 {
	var n float32

	for _, value := range answers {
		n += float32(value)
	}

	return n
}

func AnswersToArray(pos int, d map[string]int, k1 string, k2 string, k3 string, k4 string, convertArrayCh chan types.ArrayData, wg *sync.WaitGroup) {
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

	arr := types.ArrayData{
		Array: res,
		Index: pos,
	}
	convertArrayCh <- arr
}

func GetAnswers(logged bool, answer *types.Data, convertArrayCh chan types.ArrayData, wg *sync.WaitGroup) (*types.Answers, error) {

	wg.Add(4)

	go AnswersToArray(0, answer.Energy, "applianceHours", "lightBulbs", "gasTanks", "hvacHours", convertArrayCh, wg)
	go AnswersToArray(1, answer.Food, "redMeat", "whiteMeat", "dairy", "vegetarian", convertArrayCh, wg)
	go AnswersToArray(2, answer.Transport, "carKm", "publicKm", "domesticFlights", "internationalFlights", convertArrayCh, wg)
	go AnswersToArray(3, answer.Waste, "trashBags", "foodWaste", "plasticBottles", "paperPackages", convertArrayCh, wg)

	wg.Wait()
	close(convertArrayCh)

	r := make(map[int][]float32)

	for value := range convertArrayCh {
		r[value.Index] = value.Array
	}

	if len(r) != 4 {
		return nil, fmt.Errorf("there are not 4 lists, are %d", len(r))
	}

	if logged {
		err := SaveAnswersDB(r)
		if err != nil {
			log.Fatal("error saving answers in db")
		}
	} else {
		err := AnswersNotLogged(r)
		if err != nil {}
		log.Fatal("err getting answers")
	}


	return &types.Answers{
		Transport: r[2],
		Energy:    r[0],
		Waste:     r[3],
		Food:      r[1],
	}, nil

}

func AnswersNotLogged(r map[int][]float32) error {
	var wg sync.WaitGroup

	transport := types.Transport{
		CarKM:                r[2][0],
		PublicKm:             r[2][1],
		DomesticFlights:      r[2][2],
		InternationalFlights: r[2][3],
		Total:                SumAnswers(r[2]),
		User_id:              10,
	}

	food := types.Food{
		RedMeat:    r[1][0],
		WhiteMeat:  r[1][1],
		Dairy:      r[1][2],
		Vegetarian: r[1][3],
		Total:      SumAnswers(r[1]),
		User_id:    10,
	}

	waste := types.Waste{
		TrashBags:      r[3][0],
		FoodWaste:      r[3][1],
		PlasticBottles: r[3][2],
		PaperPackages:  r[3][3],
		Total:          SumAnswers(r[3]),
		User_id:        10,
	}

	energy := types.Energy{
		ApplianceHours: r[0][0],
		LightBulbs:     r[0][1],
		GasTanks:       r[0][2],
		HvacHours:      r[0][3],
		Total:          SumAnswers(r[0]),
		User_id:        10,
	}

	wg.Wait()
	
	return nil

}
