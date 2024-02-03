package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ardtieboy/lorcana_tracker/internal/card"
	"github.com/ardtieboy/lorcana_tracker/internal/persistence"
)

func main() {

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

	persistence.CreateDatabase()

	for _, card := range cards {
		err := persistence.InsertCard(card.CardID, card.Artist, card.SetID, card.SetNum, card.SetName, card.Color, card.Image, card.Cost, card.Inkable, card.Name, card.Type, card.Rarity, card.FlavorText, card.CardNum, card.BodyText, card.MarketPriceInEuro)
		if err != nil {
			fmt.Println("Error inserting card into the database:", err)
			return
		}
	}

	for _, set := range sets {
		err := persistence.InsertCardSet(set.SetID, set.SetNum, set.SetName)
		if err != nil {
			fmt.Println("Error inserting set into the database:", err)
			return
		}
	}

	// printCardsJSON(cards)
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
