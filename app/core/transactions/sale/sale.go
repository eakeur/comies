package purchase

import (
	"comies/app/core/pricing/movement"
	"comies/app/core/types"
	"math"
)

type Item struct {
	movement.Movement
	Discount float64
}

func NewItem(saleID, itemID types.ID) Item {
	return Item{
		Movement: movement.Movement{
			GroupID: saleID,
			ItemID:  itemID,
		},
	}
}

func (i Item) WithID(id types.ID) Item {
	i.ID = id
	return i
}

func (i Item) WithDiscount(discount float64) Item {
	i.Discount = discount
	return i
}

func (i Item) ValueWithDiscount() types.Currency {
	return types.Currency(math.Round(float64(i.Value) * (1 - i.Discount)))
}
