package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"encoding/json"
	"net/http"
)

func (s Service) AddToOrder(ctx context.Context, r *http.Request) handler.Response {

	var i item.Item
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	res, err := s.ordering.AddToOrder(ctx, i)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	if l := len(res.Failed); l > 0 {
		failed := make([]Failure, len(res.Failed))
		for _, f := range res.Failed {
			failed = append(failed, Failure{
				For:       f.For,
				ProductID: f.ProductID,
				Error:     f.Error,
			})
		}

		return handler.ResponseWithData(http.StatusUnprocessableEntity, failed)
	}

	return handler.ResponseWithData(http.StatusCreated, nil)
}
