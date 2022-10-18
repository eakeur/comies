package ordering

import (
	"comies/app/core/menu"
	"comies/app/core/ordering"
	"comies/app/core/types"
)

type AddToOrderRequest struct {
	ProductID      types.ID                     `json:"product_id"`
	Quantity       types.Quantity               `json:"quantity"`
	Observations   string                       `json:"observations"`
	Specifications menu.IngredientSpecification `json:"specifications"`
}

type SetItemStatusRequest struct {
	Status ordering.Status `json:"status"`
}

type SetOrderStatusRequest struct {
	Status ordering.Status `json:"status"`
}

type SetOrderDeliveryTypeRequest struct {
	Type ordering.Type `json:"mode"`
}
