package card

type Card struct {
	Artist     string `json:"Artist"`
	SetName    string `json:"Set_Name"`
	SetNum     int    `json:"Set_Num"`
	Color      string `json:"Color"`
	Image      string `json:"Image"`
	Cost       int    `json:"Cost"`
	Inkable    bool   `json:"Inkable"`
	Name       string `json:"Name"`
	Type       string `json:"Type"`
	Rarity     string `json:"Rarity"`
	FlavorText string `json:"Flavor_Text"`
	CardNum    int    `json:"Card_Num"`
	BodyText   string `json:"Body_Text"`
	SetID      string `json:"Set_ID"`
}
