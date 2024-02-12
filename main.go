package main

import (
	"github.com/ardtieboy/lorcana_tracker/controller"
	"github.com/ardtieboy/lorcana_tracker/internal/persistence"
)

func main() {
	if true {
		persistence.InitialiseState()
	}
	controller.Start()
}
