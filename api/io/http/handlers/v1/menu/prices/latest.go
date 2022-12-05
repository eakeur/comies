package prices

import (
	"comies/core/types"
	"comies/io/http/request"
	"comies/io/http/send"
	"context"
	"net/http"
)

func (h Handler) Latest(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	cur, err := h.prices.GetProductLatestPriceByID(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, GetLatestPriceResponse{Value: cur})
}

type GetLatestPriceResponse struct {
	Value types.Currency `json:"value"`
}
