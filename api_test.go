package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ardtieboy/lorcana_tracker/controller"
	"github.com/ardtieboy/lorcana_tracker/internal/card"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	router := controller.CreateRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var healthData controller.HealthData
	json.Unmarshal(w.Body.Bytes(), &healthData)
	assert.Equal(t, "OK", healthData.Health)
}

func TestGetAllCardsEndpoint(t *testing.T) {
	router := controller.CreateRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/card", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var cards []card.Card
	json.Unmarshal(w.Body.Bytes(), &cards)
	fmt.Println("Found cards: ", len(cards))
	assert.NotEmpty(t, cards)
}

func TestGetCardByIdEndpoint(t *testing.T) {
	router := controller.CreateRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/card/TFC-1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var card card.Card
	json.Unmarshal(w.Body.Bytes(), &card)
	assert.NotEmpty(t, card)
	assert.Equal(t, "TFC-1", card.CardID)
}

func TestUpdateAndGetCardInCollection(t *testing.T) {
	router := controller.CreateRouter()

	// Get 1

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/card_in_collection/TFC-1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var initialCardInCollection card.CardInCollection
	json.Unmarshal(w.Body.Bytes(), &initialCardInCollection)
	assert.Equal(t, "TFC-1", initialCardInCollection.CardID)

	// initalFoilCopies := initialCardInCollection.OwnedFoilCopies

	// Update
	w = httptest.NewRecorder()
	*initialCardInCollection.OwnedFoilCopies = 0
	*initialCardInCollection.OwnedNormalCopies = 0
	jsonData, _ := json.Marshal(initialCardInCollection)
	println(string(jsonData))
	body := bytes.NewReader(jsonData)
	req, _ = http.NewRequest("PUT", "/card_in_collection", body)
	router.ServeHTTP(w, req)
	println(w.Body.String())
	assert.Equal(t, 201, w.Code)

	// Get
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/card_in_collection/TFC-1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var cardInCollection card.CardInCollection
	json.Unmarshal(w.Body.Bytes(), &cardInCollection)
	assert.Equal(t, "TFC-1", cardInCollection.CardID)
	assert.Equal(t, 0, *cardInCollection.OwnedFoilCopies)
}
