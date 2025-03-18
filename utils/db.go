package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dsn = "host=localhost user=postgres password=root dbname=carbon_results port=5432 sslmode=disable"
var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connect db", err)
	}

}
