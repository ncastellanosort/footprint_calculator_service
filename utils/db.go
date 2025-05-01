package utils

import (
	"carbon_calculator/types"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("err getting env", e)
	}

	dsn := os.Getenv("AWS_RDS_URL")

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connect db", err)
	}

}

func SaveAnswersDB(r map[int][]float32) error {
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

	entities := []interface{}{&transport, &waste, &energy, &food}

	for _, entity := range entities {
		wg.Add(1)
		go func(e interface{}){
			defer wg.Done()
			DB.Create(e)
		}(entity)
	}

	wg.Wait()
	
	return nil

}
