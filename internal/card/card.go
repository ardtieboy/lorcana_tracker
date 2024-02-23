package card

import "fmt"

type Card struct {
	CardID     string `json:"id"`
	Artist     string `json:"artist"`
	SetID      string `json:"set_id"`
	SetNum     int    `json:"set_num"`
	SetName    string `json:"set_name"`
	Color      string `json:"color"`
	Image      string `json:"image"`
	Cost       int    `json:"cost"`
	Inkable    bool   `json:"inkable"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Rarity     string `json:"rarity"`
	FlavorText string `json:"flavor_text"`
	CardNum    int    `json:"card_num"`
	BodyText   string `json:"body_text"`
}

func (c *Card) PopulateID() {
	c.CardID = fmt.Sprintf("%s-%d", c.SetID, c.CardNum)
}
