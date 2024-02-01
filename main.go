package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"starlu.be/lorcana_tracker/internal/card"
)

func main() {
	url := "https://api.lorcana-api.com/cards/all"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var cards []card.Card
	err = json.Unmarshal(body, &cards)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, card := range cards {
		cardJSON, err := json.MarshalIndent(card, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(cardJSON))
	}
}
