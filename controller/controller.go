package controller

import (
	"net/http"

	"github.com/ardtieboy/lorcana_tracker/internal/card"
	"github.com/ardtieboy/lorcana_tracker/internal/persistence"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/cards", GetCards)
	router.PUT("/cards", UpdateCard)

	router.Run("localhost:8080")
}

func GetCards(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persistence.GetAllCards())
}

func UpdateCard(c *gin.Context) {
	var cardInCollection card.CardInCollection
	if err := c.BindJSON(&cardInCollection); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, cardInCollection)
}
