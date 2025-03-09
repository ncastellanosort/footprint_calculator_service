package main

import (
	"carbon_calculator/internal/routes"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const PORT int = 8080

func main() {

	fmt.Println(fmt.Sprintf("server running on http://localhost:%d", PORT))
	wg := &sync.WaitGroup{}

	routes.SetupCalculatorRoutes(wg)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatal(err)
	}

}
