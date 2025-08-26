package card

type InCollection struct {
	CardID            string `json:"card_id" binding:"required"`
	OwnedNormalCopies *int   `json:"owned_normal_copies" binding:"required"`
	OwnedFoilCopies   *int   `json:"owned_foil_copies" binding:"required"`
	WishList          *bool  `json:"wish_list" binding:"required"`
}
