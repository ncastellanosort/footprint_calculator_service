package calc

import (
	"bytes"
	"carbon_calculator/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func GetAnswers(logged bool, answer *types.Data, convertArrayCh chan types.ArrayData, wg *sync.WaitGroup, calculateCh chan float32) (float32, error) {
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

	a := &types.Answers{
		Transport: r[2],
		Energy:    r[0],
		Waste:     r[3],
		Food:      r[1],
	} 

	value, err := Calculator(a, calculateCh, wg)

	if logged {
		err = SaveAnswersDB(r, value)
		if err != nil {
			log.Fatal("error saving answers in db")
		}

		return value, nil

	} else {
		return value, nil
	}
}

func PostRecommendations(data types.DataRecommendation, token string) types.RecommendationResponse {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error encoding json: %v", err)
	}

	url := os.Getenv("RECOMMENDATION_URL")
	if url == "" {
		log.Printf("RECOMMENDATION_URL not set")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response body: %v", err)
	}

	var recommendation types.RecommendationResponse
	if err := json.Unmarshal(body, &recommendation); err != nil {
		fmt.Printf("error decoding JSON response: %v", err)
	}

	return recommendation 
}

