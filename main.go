package main

import (
	"flag"
	"fmt"

	"github.com/ardtieboy/lorcana_tracker/controller"
	_ "github.com/ardtieboy/lorcana_tracker/docs"
	"github.com/ardtieboy/lorcana_tracker/internal/persistence"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Lorcana Tracker API
// @version         1.0
// @description     This is a sample server for a Lorcana tracker.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8080
// @BasePath        /
func main() {
	initDB := flag.Bool("initDB", false, "Set to true if you want to initialise the database with the card data")
	flag.Parse()
	// Set to true if you want to initialise the database with the card data

	fmt.Println(*initDB)
	dbConfig := persistence.DatabaseConfig{UserDB: "lorcana_ardtieboy.db", GeneralDB: "lorcana.db"}

	if *initDB {
		err := persistence.InitialiseState(dbConfig)
		if err != nil {
			panic(err)
		}
	}
	router := controller.CreateRouter(dbConfig)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
