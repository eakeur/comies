package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) CreateProduct(ctx context.Context, p product.Product) response.Response {
	prd, err := s.menu.CreateProduct(ctx, p)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.
		WithData(http.StatusCreated, AdditionResult{ID: prd.ID})
}
