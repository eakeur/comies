package ordering

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"encoding/json"
	"net/http"
)

func (s Service) AddToOrder(ctx context.Context, r *http.Request) handler.Response {

	var i ItemAdditionRequest
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	it, err := i.ToItem(handler.GetURLParam(r, "order_id"))
	if err != nil {
		handler.IDParsingErrorResponse(err)
	}

	res, err := s.ordering.AddToOrder(ctx, it)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	if size := len(res.Failed); size > 0 {
		return handler.ResponseWithData(http.StatusUnprocessableEntity, NewFailureList(res.Failed))
	}

	return handler.ResponseWithData(http.StatusCreated, ItemAdditionResponse{ID: it.ID.String()})
}
