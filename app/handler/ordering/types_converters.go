package ordering

import (
	"comies/app/core/id"
	"comies/app/core/item"
	"comies/app/core/order"
	"comies/app/core/reservation"
	"comies/app/handler/rest"
)

func (a ItemAdditionRequest) ToItem(orderID string) (item.Item, error) {
	oid, err := rest.ConvertToID(orderID)
	if err != nil {
		return item.Item{}, err
	}

	ignore := make([]id.ID, len(a.IgnoreIngredients))
	for i, ingredient := range a.IgnoreIngredients {
		id, err := rest.ConvertToID(ingredient)
		if err != nil {
			return item.Item{}, err
		}
		ignore[i] = id
	}

	replace := make(map[id.ID]id.ID, len(a.ReplaceIngredients))
	for k, v := range a.ReplaceIngredients {
		from, err := rest.ConvertToID(k)
		if err != nil {
			return item.Item{}, err
		}

		to, err := rest.ConvertToID(v)
		if err != nil {
			return item.Item{}, err
		}

		replace[from] = to
	}

	det := item.Details{
		IgnoreIngredients:  ignore,
		ReplaceIngredients: replace,
	}

	return item.Item{
		OrderID:      oid,
		ProductID:    a.ProductID,
		Quantity:     a.Quantity,
		Observations: a.Observations,
		Details:      det,
	}, nil
}

func NewFailureList(list []reservation.Failure) []Failure {
	failures := make([]Failure, len(list))
	for _, f := range list {
		failures = append(failures, Failure{
			For:       f.For.String(),
			ProductID: f.ProductID.String(),
			Error:     f.Error,
		})
	}

	return failures
}

func NewOrder(o order.Order) Order {
	return Order{
		ID:             o.ID.String(),
		Identification: o.Identification,
		PlacedAt:       o.PlacedAt,
		Status:         o.Status,
		DeliveryMode:   o.DeliveryMode,
		Observations:   o.Observations,
		FinalPrice:     o.FinalPrice,
		Address:        o.Address,
		Phone:          o.Phone,
	}
}

func NewItemList(list []item.Item) []Item {
	items := make([]Item, len(list))
	for i, it := range list {
		items[i] = Item{
			ID:           it.ID.String(),
			OrderID:      it.OrderID.String(),
			ProductID:    it.ProductID,
			Price:        it.Price,
			Status:       it.Status,
			Quantity:     it.Quantity,
			Observations: it.Observations,
		}
	}
	return items
}

func NewOrderList(list []order.Order) []Order {
	orders := make([]Order, len(list))
	for i, o := range list {
		orders[i] = Order{
			ID:             o.ID.String(),
			Identification: o.Identification,
			PlacedAt:       o.PlacedAt,
			Status:         o.Status,
			DeliveryMode:   o.DeliveryMode,
			Observations:   o.Observations,
			FinalPrice:     o.FinalPrice,
			Address:        o.Address,
			Phone:          o.Phone,
		}
	}
	return orders
}
