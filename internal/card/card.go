package card

import "fmt"

type Card struct {
	CardID            string `json:"ID"`
	Artist            string `json:"Artist"`
	SetID             string `json:"Set_ID"`
	SetNum            int    `json:"Set_Num"`
	SetName           string `json:"Set_Name"`
	Color             string `json:"Color"`
	Image             string `json:"Image"`
	Cost              int    `json:"Cost"`
	Inkable           bool   `json:"Inkable"`
	Name              string `json:"Name"`
	Type              string `json:"Type"`
	Rarity            string `json:"Rarity"`
	FlavorText        string `json:"Flavor_Text"`
	CardNum           int    `json:"Card_Num"`
	BodyText          string `json:"Body_Text"`
	MarketPriceInEuro int    `json:"Market_Price_In_Euro"`
}

func (c *Card) PopulateID() {
	c.CardID = fmt.Sprintf("%s-%d", c.SetID, c.CardNum)
}

type CardSet struct {
	SetID   string `json:"Set_ID"`
	SetNum  int    `json:"Set_Num"`
	SetName string `json:"Set_Name"`
}

type CardInCollection struct {
	CardID            string `json:"Card_ID"`
	OwnedNormalCopies int    `json:"Owned_Normal_Copies"`
	OwnedFoilCopies   int    `json:"Owned_Foil_Copies"`
	WishLIst          bool   `json:"WishLIst"`
}

type CardView struct {
	CardID            string `json:"ID"`
	Artist            string `json:"Artist"`
	SetID             string `json:"Set_ID"`
	SetNum            int    `json:"Set_Num"`
	SetName           string `json:"Set_Name"`
	Color             string `json:"Color"`
	Image             string `json:"Image"`
	Cost              int    `json:"Cost"`
	Inkable           bool   `json:"Inkable"`
	Name              string `json:"Name"`
	Type              string `json:"Type"`
	Rarity            string `json:"Rarity"`
	FlavorText        string `json:"Flavor_Text"`
	CardNum           int    `json:"Card_Num"`
	BodyText          string `json:"Body_Text"`
	MarketPriceInEuro int    `json:"Market_Price_In_Euro"`
	OwnedNormalCopies int    `json:"Owned_Normal_Copies"`
	OwnedFoilCopies   int    `json:"Owned_Foil_Copies"`
	WishLIst          bool   `json:"WishLIst"`
}
