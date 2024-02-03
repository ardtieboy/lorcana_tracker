package persistence

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() error {
	db, err := sql.Open("sqlite3", "lorcana.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cards (card_id TEXT PRIMARY KEY UNIQUE, artist TEXT, set_id TEXT, set_num INTEGER, set_name TEXT, color TEXT, image TEXT, cost INTEGER, inkable BOOLEAN, name TEXT, card_type TEXT, rarity TEXT, flavor_text TEXT, card_num INTEGER, body_text TEXT, market_price_in_euro INTEGER)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS card_sets (set_id TEXT PRIMARY KEY UNIQUE, set_num INTEGER, set_name TEXT)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cards_in_collection (card_id TEXT PRIMARY KEY, owned_normal_copies INTEGER, owned_foil_copies INTEGER, whitelist BOOLEAN)")
	if err != nil {
		return err
	}

	log.Println("Database and tables created successfully")
	return nil
}

//card.CardID, card.Artist, card.SetID, card.SetNum, card.SetName, card.Color, card.Image, card.Cost, card.Inkable, card.Name, card.CardType, card.Rarity, card.FlavorText, card.CardNum, card.BodyText, card.MarketPriceEuro

func InsertCard(cardID string, artist string, setID string, setNum int, setName string, color string, image string, cost int, inkable bool, name string, cardType string, rarity string, flavorText string, cardNum int, bodyText string, marketPriceEuro int) error {
	db, err := sql.Open("sqlite3", "lorcana.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cards (card_id, artist, set_id, set_num, set_name, color, image, cost, inkable, name, card_type, rarity, flavor_text, card_num, body_text, market_price_in_euro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cardID, artist, setID, setNum, setName, color, image, cost, inkable, name, cardType, rarity, flavorText, cardNum, bodyText, marketPriceEuro)
	if err != nil {
		return err
	}

	log.Println("Card inserted successfully")
	return nil
}

func InsertCardSet(setID string, setNum int, setName string) error {
	db, err := sql.Open("sqlite3", "lorcana.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO card_sets (set_id, set_num, set_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(setID, setNum, setName)
	if err != nil {
		return err
	}

	log.Println("Card set inserted successfully")
	return nil
}
