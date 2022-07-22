package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) AddToOrder(ctx context.Context, it item.Item) response.Response {
	res, err := s.ordering.AddToOrder(ctx, it)
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

		return response.WithData(http.StatusUnprocessableEntity, failed)
	}

	return response.WithData(http.StatusCreated, AdditionResult{ID: it.ID})
}
