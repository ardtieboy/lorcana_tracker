package card

type CardInCollection struct {
	CardID            string `json:"card_id" binding:"required"`
	OwnedNormalCopies int    `json:"owned_normal_copies" binding:"required"`
	OwnedFoilCopies   int    `json:"owned_foil_copies" binding:"required"`
	WhishList         bool   `json:"whish_list" binding:"required"`
}
