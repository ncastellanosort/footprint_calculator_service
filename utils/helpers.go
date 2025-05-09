package utils

import (
	"bytes"
	"carbon_calculator/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
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

		return &types.Answers{
			Transport: r[2],
			Energy:    r[0],
			Waste:     r[3],
			Food:      r[1],
		}, nil

	} else {
		return &types.Answers{
			Transport: r[2],
			Energy:    r[0],
			Waste:     r[3],
			Food:      r[1],
		}, nil
	}
}

func PostData(data types.DataResponse, token string) *http.Response {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("err encoding json: %w", err)
	}

	url := os.Getenv("VUE_URL")
	if url == "" {
		log.Fatalf("err loading env")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Fatalf("err creating request %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("err making post %w", err)
	}

	return resp
}

