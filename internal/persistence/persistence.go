package persistence

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ardtieboy/lorcana_tracker/internal/card"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseConfig struct {
	UserDB    string
	GeneralDB string
}

func (dbConfig DatabaseConfig) DeleteGeneralTables() error {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE IF EXISTS cards")
	if err != nil {
		return err
	}

	_, err = db.Exec("DROP TABLE IF EXISTS card_sets")
	if err != nil {
		return err
	}

	_, err = db.Exec("DROP TABLE IF EXISTS card_prices")
	if err != nil {
		return err
	}

	log.Println("General Database and tables deleted successfully")
	return nil
}

func (dbConfig DatabaseConfig) DeleteUserTables() error {
	db, err := sql.Open("sqlite3", dbConfig.UserDB)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE IF EXISTS cards_in_collection")
	if err != nil {
		return err
	}

	log.Println("User Database and tables deleted successfully")
	return nil
}

func (dbConfig DatabaseConfig) CreateGeneralDatabaseIfNotExisting() error {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cards (card_id TEXT PRIMARY KEY UNIQUE, artist TEXT, set_id TEXT, set_num INTEGER, set_name TEXT, color TEXT, image TEXT, cost INTEGER, inkable BOOLEAN, name TEXT, card_type TEXT, rarity TEXT, flavor_text TEXT, card_num INTEGER, body_text TEXT)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS card_sets (set_id TEXT PRIMARY KEY UNIQUE, set_num INTEGER, set_name TEXT)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS card_prices (set_id TEXT PRIMARY KEY UNIQUE, market_price_in_euro INTEGER, market_price_link TEXT)")
	if err != nil {
		return err
	}

	// todo: create another table with the prices inside

	log.Println("General Database and tables created successfully")
	return nil
}

func (dbConfig DatabaseConfig) CreateUserDatabaseIfNotExisting() error {
	db, err := sql.Open("sqlite3", dbConfig.UserDB)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cards_in_collection (card_id TEXT PRIMARY KEY, owned_normal_copies INTEGER, owned_foil_copies INTEGER, whishlist BOOLEAN)")
	if err != nil {
		return err
	}

	log.Println("User Database and tables created successfully")
	return nil
}

// Cards

func (dbConfig DatabaseConfig) GetAllCards() ([]card.Card, error) {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT card_id, artist, set_id, set_num, set_name, color, image, cost, inkable, name, card_type, rarity, flavor_text, card_num, body_text from cards")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []card.Card

	for rows.Next() {
		var card card.Card

		err = rows.Scan(&card.CardID, &card.Artist, &card.SetID, &card.SetNum, &card.SetName, &card.Color, &card.Image, &card.Cost, &card.Inkable, &card.Name, &card.Type, &card.Rarity, &card.FlavorText, &card.CardNum, &card.BodyText)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func (dbConfig DatabaseConfig) GetCardById(cardID string) (card.Card, error) {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
	if err != nil {
		return card.Card{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT card_id, artist, set_id, set_num, set_name, color, image, cost, inkable, name, card_type, rarity, flavor_text, card_num, body_text from cards WHERE card_id = ?", cardID)
	if err != nil {
		return card.Card{}, err
	}
	defer rows.Close()

	var c card.Card

	if rows.Next() {
		err = rows.Scan(&c.CardID, &c.Artist, &c.SetID, &c.SetNum, &c.SetName, &c.Color, &c.Image, &c.Cost, &c.Inkable, &c.Name, &c.Type, &c.Rarity, &c.FlavorText, &c.CardNum, &c.BodyText)
		if err != nil {
			return card.Card{}, err
		}
	} else {
		return card.Card{}, errors.New("no card found with the given ID")
	}
	if err = rows.Err(); err != nil {
		return card.Card{}, err
	}

	return c, nil
}

func (dbConfig DatabaseConfig) InsertCard(c card.Card) error {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cards (card_id, artist, set_id, set_num, set_name, color, image, cost, inkable, name, card_type, rarity, flavor_text, card_num, body_text) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.CardID, c.Artist, c.SetID, c.SetNum, c.SetName, c.Color, c.Image, c.Cost, c.Inkable, c.Name, c.Type, c.Rarity, c.FlavorText, c.CardNum, c.BodyText)
	if err != nil {
		return err
	}

	log.Println("Card inserted successfully")
	return nil
}

// CardSets

func (dbConfig DatabaseConfig) GetAllCardSets() ([]card.CardSet, error) {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT set_id, set_num, set_name from card_sets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sets []card.CardSet

	for rows.Next() {
		var set card.CardSet

		err = rows.Scan(&set.SetID, &set.SetNum, &set.SetName)
		if err != nil {
			return nil, err
		}

		sets = append(sets, set)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sets, nil
}

func (dbConfig DatabaseConfig) GetCardSetById(setID string) (card.CardSet, error) {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
	if err != nil {
		return card.CardSet{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT set_id, set_num, set_name from card_sets WHERE set_id = ?", setID)
	if err != nil {
		return card.CardSet{}, err
	}
	defer rows.Close()

	var c card.CardSet

	for rows.Next() {
		err = rows.Scan(&c.SetID, &c.SetNum, &c.SetName)
		if err != nil {
			return card.CardSet{}, err
		}
	}
	if err = rows.Err(); err != nil {
		return card.CardSet{}, err
	}

	return c, nil
}

func (dbConfig DatabaseConfig) InsertCardSet(cs card.CardSet) error {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
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

// Cards in collection

func (dbConfig DatabaseConfig) GetAllCardInCollection() ([]card.CardInCollection, error) {
	db, err := sql.Open("sqlite3", dbConfig.UserDB)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT card_id, owned_normal_copies, owned_foil_copies, whishlist from cards_in_collection")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cardsInCollection []card.CardInCollection

	for rows.Next() {
		var cardInCollection card.CardInCollection

		err = rows.Scan(&cardInCollection.CardID, &cardInCollection.OwnedNormalCopies, &cardInCollection.OwnedFoilCopies, &cardInCollection.WhishList)
		if err != nil {
			return nil, err
		}

		cardsInCollection = append(cardsInCollection, cardInCollection)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cardsInCollection, nil
}

func (dbConfig DatabaseConfig) GetCardInCollectionById(cardID string) (card.CardInCollection, error) {
	db, err := sql.Open("sqlite3", dbConfig.UserDB)
	if err != nil {
		return card.CardInCollection{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT card_id, owned_normal_copies, owned_foil_copies, whishlist from cards_in_collection WHERE card_id = ?", cardID)
	if err != nil {
		return card.CardInCollection{}, err
	}
	defer rows.Close()

	var c card.CardInCollection

	if rows.Next() {
		err = rows.Scan(&c.CardID, &c.OwnedNormalCopies, &c.OwnedFoilCopies, &c.WhishList)
		if err != nil {
			return card.CardInCollection{}, err
		}
	} else {
		ownedCopies := 0
		whishList := false
		return card.CardInCollection{CardID: cardID, OwnedNormalCopies: &ownedCopies, OwnedFoilCopies: &ownedCopies, WhishList: &whishList}, nil
	}

	if err = rows.Err(); err != nil {
		return card.CardInCollection{}, err
	}
	return c, nil
}

func (dbConfig DatabaseConfig) UpdateCardInCollection(c card.CardInCollection) error {
	db, err := sql.Open("sqlite3", dbConfig.UserDB)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cards_in_collection (card_id, owned_normal_copies, owned_foil_copies, whishlist) VALUES (?, ?, ?, ?) ON CONFLICT(card_id) DO UPDATE SET owned_normal_copies=excluded.owned_normal_copies, owned_foil_copies=excluded.owned_foil_copies, whishlist=excluded.whishlist")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.CardID, c.OwnedNormalCopies, c.OwnedFoilCopies, c.WhishList)
	if err != nil {
		return err
	}

	log.Println("Card in collection updated successfully")
	return nil
}

// CardSets

func (dbConfig DatabaseConfig) GetCardPriceById(setID string) (card.CardPrice, error) {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
	if err != nil {
		return card.CardPrice{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT set_id, market_price_in_euro, market_price_link from card_prices WHERE set_id = ?", setID)
	if err != nil {
		return card.CardPrice{}, err
	}
	defer rows.Close()

	var c card.CardPrice

	for rows.Next() {
		err = rows.Scan(&c.CardID, &c.MarketPriceInEuro, &c.MarketPriceLink)
		if err != nil {
			return card.CardPrice{}, err
		}
	}
	if err = rows.Err(); err != nil {
		return card.CardPrice{}, err
	}

	return c, nil
}

func (dbConfig DatabaseConfig) UpdateCardPrice(cp card.CardPrice) error {
	db, err := sql.Open("sqlite3", dbConfig.GeneralDB)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO card_prices (set_id, market_price_in_euro, market_price_link) VALUES (?, ?, ?) ON CONFLICT(set_id) DO UPDATE SET market_price_in_euro=excluded.market_price_in_euro, market_price_link=excluded.market_price_link")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cp.CardID, cp.MarketPriceInEuro, cp.MarketPriceLink)
	if err != nil {
		return err
	}

	log.Println("Card price updated successfully")
	return nil
}
