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

var databaseConfig persistence.DatabaseConfig

func CreateRouter(dbConfig persistence.DatabaseConfig) *gin.Engine {
	router := gin.Default()
	databaseConfig = dbConfig

	router.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, HealthData{Health: "OK"})
	})

	router.GET("/card", GetAllCards)
	router.GET("/card/:id", GetCardById)

	router.GET("/card_in_collection", GetAllCardsInCollectionById)
	router.GET("/card_in_collection/:id", GetCardInCollectionById)
	router.PUT("/card_in_collection", UpdateCardInCollection)

	router.GET("/set", GetAllSets)

	router.GET("card_price/:id", GetCardPriceById)

	return router
}

// GetAllCards godoc
// @Summary Get all cards
// @Description Get all cards
// @Produce  json
// @Success 200 {array} card.Card
// @Failure 500 {string} string "Internal Server Error"
// @Router /card [get]
func GetAllCards(c *gin.Context) {
	fetchedCards, err := databaseConfig.GetAllCards()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedCards)
	}
}

// GetCardById godoc
// @Summary Get a card by ID
// @Description Get a card by ID
// @Produce  json
// @Param id path string true "Card ID"
// @Success 200 {object} card.Card
// @Failure 404 {string} string "no card found with the given ID"
// @Failure 500 {string} string "Internal Server Error"
// @Router /card/{id} [get]
func GetCardById(c *gin.Context) {
	cardId := c.Param("id")
	fetchedCard, err := databaseConfig.GetCardById(cardId)
	if err != nil {
		if err.Error() == "no card found with the given ID" {
			c.IndentedJSON(http.StatusNotFound, err.Error())
		} else {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.IndentedJSON(http.StatusOK, fetchedCard)
	}
}

// GetCardInCollectionById godoc
// @Summary Get a card in collection by ID
// @Description Get a card in collection by ID
// @Produce  json
// @Param id path string true "Card ID"
// @Success 200 {object} card.InCollection
// @Failure 500 {string} string "Internal Server Error"
// @Router /card_in_collection/{id} [get]
func GetCardInCollectionById(c *gin.Context) {
	cardId := c.Param("id")
	fetchedCard, err := databaseConfig.GetCardInCollectionById(cardId)
	fmt.Println("Fetched card: ", fetchedCard)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedCard)
	}
}

// GetAllCardsInCollectionById godoc
// @Summary Get all cards in collection
// @Description Get all cards in collection
// @Produce  json
// @Success 200 {array} card.InCollection
// @Failure 500 {string} string "Internal Server Error"
// @Router /card_in_collection [get]
func GetAllCardsInCollectionById(c *gin.Context) {
	fetchedCards, err := databaseConfig.GetAllCardInCollection()
	fmt.Println("Fetched cards: ", fetchedCards)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedCards)
	}
}

// UpdateCardInCollection godoc
// @Summary Update a card in collection
// @Description Update a card in collection
// @Accept  json
// @Produce  json
// @Param cardInCollection body card.InCollection true "Card in collection"
// @Success 201 {object} card.InCollection
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /card_in_collection [put]
func UpdateCardInCollection(c *gin.Context) {
	var cardInCollection card.InCollection
	if err := c.BindJSON(&cardInCollection); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := databaseConfig.UpdateCardInCollection(cardInCollection)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusCreated, cardInCollection)
	}
}

// GetAllSets godoc
// @Summary Get all sets
// @Description Get all sets
// @Produce  json
// @Success 200 {array} card.Set
// @Failure 500 {string} string "Internal Server Error"
// @Router /set [get]
func GetAllSets(c *gin.Context) {
	fetchedSets, err := databaseConfig.GetAllCardSets()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedSets)
	}
}

// GetCardPriceById godoc
// @Summary Get a card price by ID
// @Description Get a card price by ID
// @Produce  json
// @Param id path string true "Card ID"
// @Success 200 {object} card.Price
// @Failure 500 {string} string "Internal Server Error"
// @Router /card_price/{id} [get]
func GetCardPriceById(c *gin.Context) {
	setId := c.Param("id")
	fetchedCardPrice, err := databaseConfig.GetCardPriceById(setId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedCardPrice)
	}
}
