package controller

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/ardtieboy/lorcana_tracker/internal/card"
	"github.com/ardtieboy/lorcana_tracker/internal/persistence"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/cards", GetAllCards)
	router.PUT("/cards", UpdateCard)

	router.GET("/sets", GetAllSets)

	router.Run("localhost:8080")
}

func GetAllSets(c *gin.Context) {
	fetchedSets, err := persistence.GetAllSets()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, fetchedSets)
	}
}

func GetAllCards(c *gin.Context) {
	fetchedCards, err := persistence.GetAllCards()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}
	var filteredFetchedCards []card.CardView

	setIDs, present := c.GetQueryArray("set")
	println("SetIDs:", len(setIDs))
	collection, present := c.GetQuery("collection")
	if present {
		if collection != "all" && collection != "owned" && collection != "missing" {
			c.IndentedJSON(http.StatusBadRequest, "Invalid value for collection")
			return
		}
		filteredFetchedCards, err = postProcessCards(fetchedCards, collection, setIDs)
	} else {
		filteredFetchedCards, err = postProcessCards(fetchedCards, "all", setIDs)
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, filteredFetchedCards)
	}
}

func postProcessCards(cards []card.CardView, collection string, setIDs []string) ([]card.CardView, error) {
	var cardViewsAfterCollectionFilter []card.CardView
	for _, c := range cards {
		if collection == "all" {
			cardViewsAfterCollectionFilter = append(cardViewsAfterCollectionFilter, c)
		} else if collection == "owned" {
			if c.OwnedNormalCopies > 0 || c.OwnedFoilCopies > 0 {
				cardViewsAfterCollectionFilter = append(cardViewsAfterCollectionFilter, c)
			}
		} else if collection == "missing" {
			if c.OwnedNormalCopies == 0 && c.OwnedFoilCopies == 0 {
				cardViewsAfterCollectionFilter = append(cardViewsAfterCollectionFilter, c)
			}
		} else {
			return nil, errors.New("Invalid value for collection")
		}
	}
	fmt.Println("Size of cardViews after filtering by collection:", len(cardViewsAfterCollectionFilter))

	var cardViewsAfterSetFilter []card.CardView = []card.CardView{}
	if len(setIDs) == 0 {
		cardViewsAfterSetFilter = cardViewsAfterCollectionFilter
	} else {
		for _, c := range cardViewsAfterCollectionFilter {
			if slices.Contains(setIDs, c.SetID) {
				cardViewsAfterSetFilter = append(cardViewsAfterSetFilter, c)
			}
		}
	}
	fmt.Println("Size of cardViews after filtering by set:", len(cardViewsAfterSetFilter))
	return cardViewsAfterSetFilter, nil

}

func UpdateCard(c *gin.Context) {
	var cardInCollection card.CardInCollection
	if err := c.BindJSON(&cardInCollection); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	updatedCardId, err := persistence.UpdateCardInCollection(cardInCollection)
	fmt.Printf("Updated card with ID %s\n", updatedCardId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusCreated, cardInCollection)
	}
}
