package persistence

import (
	"database/sql"
	"log"

	"github.com/ardtieboy/lorcana_tracker/internal/card"
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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cards_in_collection (card_id TEXT PRIMARY KEY, owned_normal_copies INTEGER, owned_foil_copies INTEGER, whishlist BOOLEAN)")
	if err != nil {
		return err
	}

	// create another table with the prices inside

	log.Println("Database and tables created successfully")
	return nil
}

func InsertCard(c card.Card) error {
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

	_, err = stmt.Exec(c.CardID, c.Artist, c.SetID, c.SetNum, c.SetName, c.Color, c.Image, c.Cost, c.Inkable, c.Name, c.Type, c.Rarity, c.FlavorText, c.CardNum, c.BodyText, c.MarketPriceInEuro)
	if err != nil {
		return err
	}

	log.Println("Card inserted successfully")
	return nil
}

func InsertCardSet(cs card.CardSet) error {
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

	_, err = stmt.Exec(cs.SetID, cs.SetNum, cs.SetName)
	if err != nil {
		return err
	}

	log.Println("Card set inserted successfully")
	return nil
}

func GetAllCards() []card.CardView {
	db, err := sql.Open("sqlite3", "lorcana.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT cards.card_id, artist, set_id, set_num, set_name, color, image, cost, inkable,name, card_type, rarity, flavor_text, card_num, body_text, market_price_in_euro, owned_normal_copies, owned_foil_copies, whishlist from cards left join cards_in_collection on cards.card_id=cards_in_collection.card_id ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cards []card.CardView

	for rows.Next() {
		var view CardDatabaseView

		err = rows.Scan(&view.CardID, &view.Artist, &view.SetID, &view.SetNum, &view.SetName, &view.Color, &view.Image, &view.Cost, &view.Inkable, &view.Name, &view.Type, &view.Rarity, &view.FlavorText, &view.CardNum, &view.BodyText, &view.MarketPriceInEuro, &view.OwnedNormalCopies, &view.OwnedFoilCopies, &view.WhishList)
		if err != nil {
			log.Fatal(err)
		}

		cards = append(cards, view.toCardView())

		// Process the card data here
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return cards

}
