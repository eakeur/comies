package menu

import (
	"comies/api/request"
	"comies/api/send"
	"comies/core/types"
	"context"
	"net/http"
)

func (h Handler) GetLatestItemPrice(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam(ItemIDParam)
	if err != nil {
		return send.IDError(err)
	}

	cur, err := h.menu.GetProductLatestPriceByID(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, GetLatestPriceResponse{Value: cur})
}

type GetLatestPriceResponse struct {
	Value types.Currency `json:"value"`
}
