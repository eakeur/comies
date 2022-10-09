package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/types"
	"comies/app/gateway/api/handler"
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

// ListItems
//
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Success     200         {object} handler.Response{data=[]Item{}}
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/items [GET]
func (s Service) ListItems(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	items, err := s.ordering.ListItems(ctx, id)
	if err != nil {
		return handler.Fail(err)
	}

	return handler.ResponseWithData(http.StatusOK, NewItemList(items))

}
