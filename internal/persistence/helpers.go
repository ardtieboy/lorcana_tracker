package persistence

import (
	"database/sql"

	"github.com/ardtieboy/lorcana_tracker/internal/card"
)

type CardDatabaseView struct {
	CardID            string        `json:"ID"`
	Artist            string        `json:"Artist"`
	SetID             string        `json:"Set_ID"`
	SetNum            int           `json:"Set_Num"`
	SetName           string        `json:"Set_Name"`
	Color             string        `json:"Color"`
	Image             string        `json:"Image"`
	Cost              int           `json:"Cost"`
	Inkable           bool          `json:"Inkable"`
	Name              string        `json:"Name"`
	Type              string        `json:"Type"`
	Rarity            string        `json:"Rarity"`
	FlavorText        string        `json:"Flavor_Text"`
	CardNum           int           `json:"Card_Num"`
	BodyText          string        `json:"Body_Text"`
	MarketPriceInEuro int           `json:"Market_Price_In_Euro"`
	OwnedNormalCopies sql.NullInt32 `json:"Owned_Normal_Copies"`
	OwnedFoilCopies   sql.NullInt32 `json:"Owned_Foil_Copies"`
	WhishList         sql.NullBool  `json:"WhishList"`
}

func (c *CardDatabaseView) toCardView() card.CardView {

	var ownedNormalCopies int
	if c.OwnedNormalCopies.Valid {
		ownedNormalCopies = int(c.OwnedNormalCopies.Int32)

	} else {
		ownedNormalCopies = 0
	}

	var ownedFoilCopies int
	if c.OwnedFoilCopies.Valid {
		ownedFoilCopies = int(c.OwnedFoilCopies.Int32)

	} else {
		ownedFoilCopies = 0
	}

	var whishList bool
	if c.WhishList.Valid {
		whishList = c.WhishList.Bool

	} else {
		whishList = false
	}

	return card.CardView{
		CardID: c.CardID, Artist: c.Artist, SetID: c.SetID, SetNum: c.SetNum, SetName: c.SetName, Color: c.Color, Image: c.Image,
		Cost: c.Cost, Inkable: c.Inkable, Name: c.Name, Type: c.Type, Rarity: c.Rarity, FlavorText: c.FlavorText, CardNum: c.CardNum,
		BodyText: c.BodyText, MarketPriceInEuro: c.MarketPriceInEuro,
		OwnedNormalCopies: ownedFoilCopies, OwnedFoilCopies: ownedFoilCopies, WhishList: whishList}
}
