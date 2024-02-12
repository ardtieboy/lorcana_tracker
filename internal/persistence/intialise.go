package persistence

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ardtieboy/lorcana_tracker/internal/card"
)

func InitialiseState() {
	fmt.Println("Fetching cards from API")

	cards, err := fetchAllCardsFromApi()
	if err != nil {
		fmt.Println("Error fetching cards from API:", err)
		return
	}

	sets, err := extractSetsFromCards(cards)
	if err != nil {
		fmt.Println("Error extracting sets from cards:", err)
		return
	}

	fmt.Println("Size of cards:", len(cards))
	fmt.Println("Size of sets:", len(sets))

	CreateDatabase()

	for _, c := range cards {
		err := InsertCard(*c)
		if err != nil && err.Error() == "UNIQUE constraint failed: cards.card_id" {
			fmt.Printf("Ignoring insertion of card with ID %s because it already exists in the database.\n", c.CardID)
		} else if err != nil {
			fmt.Println("Error inserting card into the database:", err)
			return
		}
	}

	for _, set := range sets {
		err := InsertCardSet(set)
		if err != nil && err.Error() == "UNIQUE constraint failed: card_sets.set_id" {
			fmt.Printf("Ignoring insertion of set with ID %s because it already exists in the database.\n", set.SetID)
		} else if err != nil {
			fmt.Println("Error inserting set into the database:", err)
			return
		}
	}

	fmt.Println("========================")
	fmt.Println("========================")
	fmt.Println("========================")
}

func extractSetsFromCards(cards []*card.Card) ([]card.CardSet, error) {
	setIDs := make(map[string]card.CardSet)

	for _, c := range cards {
		setIDs[c.SetID] = card.CardSet{SetID: c.SetID, SetNum: c.SetNum, SetName: c.SetName}
	}

	var sets []card.CardSet
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

	for _, card := range cards {
		card.PopulateID()
	}
	return cards, nil
}

func printCardsJSON(cards []*card.Card) {
	jsonData, err := json.MarshalIndent(cards, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling cards to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
