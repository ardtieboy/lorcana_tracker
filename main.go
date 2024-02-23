package main

import (
	"flag"
	"fmt"

	"github.com/ardtieboy/lorcana_tracker/controller"
	"github.com/ardtieboy/lorcana_tracker/internal/persistence"
)

func main() {
	initDB := flag.Bool("initDB", false, "Set to true if you want to initialise the database with the card data")
	flag.Parse()
	// Set to true if you want to initialise the database with the card data

	fmt.Println(*initDB)

	if *initDB {
		err := persistence.InitialiseState()
		if err != nil {
			panic(err)
		}
	}
	router := controller.CreateRouter()
	router.Run(":8080")
}
