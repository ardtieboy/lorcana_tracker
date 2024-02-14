package main

import (
	"github.com/ardtieboy/lorcana_tracker/controller"
	"github.com/ardtieboy/lorcana_tracker/internal/persistence"
)

func main() {
	// Set to true if you want to initialise the database with the card data
	if false {
		err := persistence.InitialiseState()
		if err != nil {
			panic(err)
		}
	}
	controller.Start()
}
