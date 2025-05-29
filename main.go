package main

import (
	"carbon_calculator/internal"
	"carbon_calculator/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("assuming production env")
		}
	}

	utils.InitDB()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatalf("port not set in env")
	}
	fmt.Printf("server running on http://localhost:%s\n", PORT)

	wg := &sync.WaitGroup{}
	internal.SetupCalculatorRoutes(wg)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil); err != nil {
		log.Fatalf("err starting %v", err)
	}

}
