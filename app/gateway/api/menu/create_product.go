package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) CreateProduct(ctx context.Context, p Product) response.Response {
	prd, err := s.menu.CreateProduct(ctx, product.Product{
		Code: p.Code,
		Name: p.Name,
		Type: p.Type,
		Sale: product.Sale{
			CostPrice:   p.CostPrice,
			SalePrice:   p.SalePrice,
			SaleUnit:    p.SaleUnit,
			MinimumSale: p.MinimumSale,
		},
		Stock: product.Stock{
			MaximumQuantity: p.MaximumQuantity,
			MinimumQuantity: p.MinimumQuantity,
			Location:        p.Location,
		},
	})
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.
		WithData(http.StatusCreated, AdditionResult{ID: prd.ID.String()})
}
