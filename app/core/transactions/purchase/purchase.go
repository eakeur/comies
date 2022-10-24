package purchase

import (
	"comies/app/core/pricing/movement"
	"comies/app/core/types"
)

type Item = movement.Movement

func NewItem(purchaseID, itemID types.ID) Item {
	return movement.Movement{
		GroupID: purchaseID,
		ItemID:  itemID,
	}
}
