package ordering

import (
	"comies/app/core/id"
	"comies/app/core/item"
	"comies/app/core/types"
	"comies/app/core/workflows/ordering"
	"comies/app/handler/rest"
	"context"
	"net/http"
)

type ListItemsResponse struct {
	ID                 id.ID           `json:"id,omitempty"`
	OrderID            id.ID           `json:"order_id"`
	Status             item.Status     `json:"status"`
	Price              types.Currency  `json:"price"`
	ProductID          id.ID           `json:"product_id"`
	Quantity           types.Quantity  `json:"quantity"`
	Observations       string          `json:"observations"`
	IgnoreIngredients  []id.ID         `json:"ignore_ingredients"`
	ReplaceIngredients map[id.ID]id.ID `json:"replace_ingredients"`
}

// ListItems
//
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Success     200         {object} rest.Response{data=[]Item{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/items [GET]
func ListItems(ctx context.Context, r *http.Request) rest.Response {
	id, err := rest.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	items, err := ordering.ListItems(ctx, id)
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusOK, NewItemList(items))

}
