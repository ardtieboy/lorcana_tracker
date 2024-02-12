package card

type CardInCollection struct {
	CardID            string `json:"card_id"`
	OwnedNormalCopies int    `json:"owned_normal_copies"`
	OwnedFoilCopies   int    `json:"owned_foil_copies"`
	WhishList         bool   `json:"whish_list"`
}
