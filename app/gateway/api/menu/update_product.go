package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) UpdateProduct(ctx context.Context, p Product) response.Response {
	productID, e, res := convertToID(p.ID)
	if e != nil {
		return res
	}

	err := s.menu.UpdateProduct(ctx, product.Product{
		ID:   productID,
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

	return response.WithData(http.StatusNoContent, nil)
}
