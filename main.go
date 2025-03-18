package main

import (
	"carbon_calculator/internal"
	"carbon_calculator/types"
	"carbon_calculator/utils"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const PORT int = 8080

func main() {

	utils.Connect()

	err := utils.DB.AutoMigrate(&types.Transport{}, &types.Energy{}, &types.Waste{}, &types.Food{})

	if err != nil {
		log.Fatal("Migrate err", err)
	}
	fmt.Println(fmt.Sprintf("server running on http://localhost:%d", PORT))
	wg := &sync.WaitGroup{}

	internal.SetupCalculatorRoutes(wg)

	err = http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatal(err)
	}

}
