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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cards_in_collection (card_id TEXT PRIMARY KEY, owned_normal_copies INTEGER, owned_foil_copies INTEGER, whitelist BOOLEAN)")
	if err != nil {
		return err
	}

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
	if err != nil && err.Error() == "UNIQUE constraint failed: cards.card_id" {
		log.Println("Ignoring the insert of " + c.CardID + " as the card is already in the database")
		return nil
	} else if err != nil {
		return err
	} else {
		log.Println("Card " + c.CardID + " inserted successfully")
		return nil
	}

}

func InsertCardSet(cardSet card.CardSet) error {
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

	_, err = stmt.Exec(cardSet.SetID, cardSet.SetNum, cardSet.SetName)
	if err != nil && err.Error() == "UNIQUE constraint failed: card_sets.set_id" {
		log.Println("Ignoring the insert of " + cardSet.SetID + " as the set is already in the database")
		return nil
	} else if err != nil {
		return err
	} else {
		log.Println("Set " + cardSet.SetID + " inserted successfully")
		return nil
	}
}

func FetchAllCardsFromDatabase() ([]*card.CardView, error) {
	db, err := sql.Open("sqlite3", "lorcana.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select cards.card_id, artist, set_id, set_num, set_name, color, image, cost, inkable, name, card_type, rarity, flavor_text, card_num, body_text, market_price_in_euro, owned_normal_copies, owned_foil_copies, whitelist from cards left join cards_in_collection on cards.card_id=cards_in_collection.card_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cardViews []*card.CardView
	for rows.Next() {
		var cardView card.CardView
		// TODO: Cards for which there is no entry in the cards_in_collection table will have NULL values for the columns in that table --> ERROR
		err = rows.Scan(&cardView.CardID, &cardView.Artist, &cardView.SetID, &cardView.SetNum, &cardView.SetName, &cardView.Color, &cardView.Image, &cardView.Cost, &cardView.Inkable, &cardView.Name, &cardView.Type, &cardView.Rarity, &cardView.FlavorText, &cardView.CardNum, &cardView.BodyText, &cardView.MarketPriceInEuro, &cardView.OwnedNormalCopies, &cardView.OwnedFoilCopies, &cardView.WishLIst)
		if err != nil {
			return nil, err
		}
		cardViews = append(cardViews, &cardView)
	}

	return cardViews, nil
}
