package main

import (
	"carbon_calculator/internal"
	"carbon_calculator/types"
	"carbon_calculator/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {

	if e := godotenv.Load(); e != nil {
		log.Fatalf("err getting envs %v", e)
	}

	utils.Connect()

	if err := utils.DB.AutoMigrate(
		&types.Transport{},
		&types.Energy{},
		&types.Waste{},
		&types.Food{},
	); err != nil {
		log.Fatalf("migrate err %v", err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatalf("port not set in env")
	}
	fmt.Printf("server running on http://localhost:%s", PORT)

	wg := &sync.WaitGroup{}
	internal.SetupCalculatorRoutes(wg)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil); err != nil {
		log.Fatalf("err starting %v", err)
	}

}
