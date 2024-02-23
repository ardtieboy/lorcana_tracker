package card

type CardPrice struct {
	CardID            string `json:"card_id" binding:"required"`
	MarketPriceInEuro int    `json:"market_price_in_euro"`
	MarketPriceLink   string `json:"market_price_link"`
}
