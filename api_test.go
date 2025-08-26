package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ardtieboy/lorcana_tracker/controller"
	"github.com/ardtieboy/lorcana_tracker/internal/card"
	"github.com/ardtieboy/lorcana_tracker/internal/persistence"
	"github.com/stretchr/testify/assert"
)

func InitialiseTestState(dbConfig persistence.DatabaseConfig) error {
	fmt.Println("Fetching cards from API")

	cards := []*card.Card{{
		CardID:  "AAA-1",
		Artist:  "Dummy Artist",
		SetID:   "AAA",
		SetNum:  1,
		SetName: "Dummy Set Name",
		Color:   "Dummy Color",
		Image:   "Dummy Image",
		Cost:    1,
		Inkable: true,
		Name:    "Dummy Name",
		Type:    "Dummy Type",
		Rarity:  "Dummy Rarity",
	}}

	sets := []card.Set{{SetID: "AAA", SetNum: 1, SetName: "Dummy Set Name"}}

	fmt.Println("Size of cards provided by the api:", len(cards))
	fmt.Println("Size of sets provied by the api:", len(sets))

	dbConfig.DeleteGeneralTables()
	dbConfig.DeleteUserTables()
	dbConfig.CreateGeneralDatabaseIfNotExisting()
	dbConfig.CreateUserDatabaseIfNotExisting()

	newlyInsertedCards := 0
	newlyInsertedSets := 0

	for _, c := range cards {
		err := dbConfig.InsertCard(*c)
		if err != nil && err.Error() == "UNIQUE constraint failed: cards.card_id" {
			fmt.Printf("Ignoring insertion of card with ID %s because it already exists in the database.\n", c.CardID)
		} else if err != nil {
			fmt.Println("Error inserting card into the database:", err)
			return err
		} else {
			newlyInsertedCards++
		}
	}

	for _, set := range sets {
		err := dbConfig.InsertCardSet(set)
		if err != nil && err.Error() == "UNIQUE constraint failed: card_sets.set_id" {
			fmt.Printf("Ignoring insertion of set with ID %s because it already exists in the database.\n", set.SetID)
		} else if err != nil {
			fmt.Println("Error inserting set into the database:", err)
			return err
		} else {
			newlyInsertedSets++
		}
	}

	fmt.Println("========================")
	fmt.Println("========================")
	fmt.Printf("Inserted %d new cards into the database\n", newlyInsertedCards)
	fmt.Printf("Inserted %d new sets into the database\n", newlyInsertedSets)
	fmt.Println("========================")
	fmt.Println("========================")
	return nil
}

var testDatabaseConfig = persistence.DatabaseConfig{UserDB: "test_ardtieboy.db", GeneralDB: "test.db"}

// TestMain is always run when the tests are being run

func TestMain(m *testing.M) {
	InitialiseTestState(testDatabaseConfig)
	code := m.Run()
	os.Exit(code)
}

// The real testing begins here...

func TestHealthEndpoint(t *testing.T) {
	router := controller.CreateRouter(testDatabaseConfig)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var healthData controller.HealthData
	json.Unmarshal(w.Body.Bytes(), &healthData)
	assert.Equal(t, "OK", healthData.Health)
}

func TestGetAllCardsEndpoint(t *testing.T) {
	router := controller.CreateRouter(testDatabaseConfig)

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
	router := controller.CreateRouter(testDatabaseConfig)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/card/AAA-1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var card card.Card
	json.Unmarshal(w.Body.Bytes(), &card)
	assert.NotEmpty(t, card)
	assert.Equal(t, "AAA-1", card.CardID)
}

func TestGetCardByIdEndpointNotExisting(t *testing.T) {
	router := controller.CreateRouter(testDatabaseConfig)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/card/BBB-1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestUpdateAndGetCardInCollection(t *testing.T) {
	router := controller.CreateRouter(testDatabaseConfig)

	// Get 1
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/card_in_collection/AAA-1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var initialCardInCollection card.InCollection
	json.Unmarshal(w.Body.Bytes(), &initialCardInCollection)
	assert.Equal(t, "AAA-1", initialCardInCollection.CardID)

	fmt.Println("Initial card in collection: ", initialCardInCollection)

	initalFoilCopies := *initialCardInCollection.OwnedFoilCopies
	initalFoilCopies++

	// Update
	w = httptest.NewRecorder()
	initialCardInCollection.OwnedFoilCopies = &initalFoilCopies
	jsonData, _ := json.Marshal(initialCardInCollection)
	println(string(jsonData))
	body := bytes.NewReader(jsonData)
	req, _ = http.NewRequest("PUT", "/card_in_collection", body)
	router.ServeHTTP(w, req)
	println(w.Body.String())
	assert.Equal(t, 201, w.Code)

	// Get
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/card_in_collection/AAA-1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var cardInCollection card.InCollection
	json.Unmarshal(w.Body.Bytes(), &cardInCollection)
	assert.Equal(t, "AAA-1", cardInCollection.CardID)
	assert.Equal(t, initalFoilCopies, *cardInCollection.OwnedFoilCopies)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/card_in_collection", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var cardsInCollection []card.InCollection
	json.Unmarshal(w.Body.Bytes(), &cardsInCollection)
	//assert.Equal(t, "AAA-1", cardsInCollection[0].CardID)
	//assert.Equal(t, initalFoilCopies, *cardsInCollection[0].OwnedFoilCopies)

	initalFoilCopies--

	// Put the card back to the initial state
	w = httptest.NewRecorder()
	*initialCardInCollection.OwnedFoilCopies = initalFoilCopies
	jsonData, _ = json.Marshal(initialCardInCollection)
	println(string(jsonData))
	body = bytes.NewReader(jsonData)
	req, _ = http.NewRequest("PUT", "/card_in_collection", body)
	router.ServeHTTP(w, req)
	println(w.Body.String())
	assert.Equal(t, 201, w.Code)
}
