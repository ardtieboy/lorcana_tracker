package card

type CardView struct {
	CardID            string `json:"id"`
	Artist            string `json:"artist"`
	SetID             string `json:"set_id"`
	SetNum            int    `json:"set_num"`
	SetName           string `json:"set_name"`
	Color             string `json:"color"`
	Image             string `json:"image"`
	Cost              int    `json:"cost"`
	Inkable           bool   `json:"inkable"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Rarity            string `json:"rarity"`
	FlavorText        string `json:"flavor_text"`
	CardNum           int    `json:"card_num"`
	BodyText          string `json:"body_text"`
	MarketPriceInEuro int    `json:"market_price_in_euro"`
	OwnedNormalCopies int    `json:"owned_normal_copies"`
	OwnedFoilCopies   int    `json:"owned_foil_copies"`
	WhishList         bool   `json:"whish_list"`
}
