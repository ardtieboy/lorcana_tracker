package controller

import (
	"fmt"
	"net/http"

	"github.com/ardtieboy/lorcana_tracker/internal/card"
	"github.com/ardtieboy/lorcana_tracker/internal/persistence"
	"github.com/gin-gonic/gin"
)

type HealthData struct {
	Health string `json:"Health"`
}

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, HealthData{Health: "OK"})
	})

	router.GET("/card", GetAllCards)
	router.GET("/card/:id", GetCardById)

	router.GET("/card_in_collection/:id", GetCardInCollectionById)
	router.PUT("/card_in_collection", UpdateCardInCollection)

	router.GET("/set", GetAllSets)

	router.GET("card_price/:id", GetCardPriceById)

	return router
}

func GetAllCards(c *gin.Context) {
	fetchedCards, err := persistence.GetAllCards()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedCards)
	}
}

func GetCardById(c *gin.Context) {
	cardId := c.Param("id")
	fetchedCard, err := persistence.GetCardById(cardId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedCard)
	}
}

func GetCardInCollectionById(c *gin.Context) {
	cardId := c.Param("id")
	fetchedCard, err := persistence.GetCardInCollectionById(cardId)
	fmt.Println("Fetched card: ", fetchedCard)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedCard)
	}
}

func UpdateCardInCollection(c *gin.Context) {
	var cardInCollection card.CardInCollection
	if err := c.BindJSON(&cardInCollection); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := persistence.UpdateCardInCollection(cardInCollection)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusCreated, cardInCollection)
	}
}

func GetAllSets(c *gin.Context) {
	fetchedSets, err := persistence.GetAllCardSets()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedSets)
	}
}

func GetCardPriceById(c *gin.Context) {
	setId := c.Param("id")
	fetchedCardPrice, err := persistence.GetCardPriceById(setId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedCardPrice)
	}
}
