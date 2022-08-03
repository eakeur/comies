package ordering

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) ListItems(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return res
	}

	items, err := s.ordering.ListItems(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	list := make([]Item, len(items))
	for i, it := range items {
		list[i] = Item{
			ID:           it.ID,
			OrderID:      it.OrderID,
			ProductID:    it.ProductID,
			Price:        it.Price,
			Status:       it.Status,
			Quantity:     it.Quantity,
			Observations: it.Observations,
		}
	}

	return handler.ResponseWithData(http.StatusOK, list)

}
