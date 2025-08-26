package persistence

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ardtieboy/lorcana_tracker/internal/card"
)

func InitialiseState(dbConfig DatabaseConfig) error {
	fmt.Println("Fetching cards from API")

	cards, err := fetchAllCardsFromApi()
	if err != nil {
		fmt.Println("Error fetching cards from API: ", err)
		return err
	}

	sets, err := extractSetsFromCards(cards)
	if err != nil {
		fmt.Println("Error extracting sets from cards:", err)
		return err
	}

	fmt.Println("Size of cards provided by the api:", len(cards))
	fmt.Println("Size of sets provied by the api:", len(sets))

	err = dbConfig.CreateGeneralDatabaseIfNotExisting()
	if err != nil {
		return err
	}
	err = dbConfig.CreateUserDatabaseIfNotExisting()
	if err != nil {
		return err
	}

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

func extractSetsFromCards(cards []*card.Card) ([]card.Set, error) {
	setIDs := make(map[string]card.Set)

	for _, c := range cards {
		setIDs[c.SetID] = card.Set{SetID: c.SetID, SetNum: c.SetNum, SetName: c.SetName}
	}

	var sets []card.Set
	for _, value := range setIDs {
		sets = append(sets, value)
	}

	return sets, nil
}

func fetchAllCardsFromApi() ([]*card.Card, error) {
	url := "https://api.lorcana-api.com/cards/all"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return nil, err
	}

	var cards []*card.Card
	err = json.Unmarshal(body, &cards)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Create the card IDs here, because the API does not provide them
	for _, card := range cards {
		card.PopulateID()
	}
	return cards, nil
}
