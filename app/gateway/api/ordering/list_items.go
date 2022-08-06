package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"net/http"
)

type ListItemsResponse struct {
	ID                 types.ID              `json:"id,omitempty"`
	OrderID            types.ID              `json:"order_id"`
	Status             item.Status           `json:"status"`
	Price              types.Currency        `json:"price"`
	ProductID          types.ID              `json:"product_id"`
	Quantity           types.Quantity        `json:"quantity"`
	Observations       string                `json:"observations"`
	IgnoreIngredients  []types.ID            `json:"ignore_ingredients"`
	ReplaceIngredients map[types.ID]types.ID `json:"replace_ingredients"`
}

func (s Service) ListItems(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	items, err := s.ordering.ListItems(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusOK, NewItemList(items))

}
