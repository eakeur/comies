package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"encoding/json"
	"net/http"
)

func (s Service) UpdateProduct(ctx context.Context, r *http.Request) handler.Response {

	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	productID, e, res := handler.ConvertToID(p.ID)
	if e != nil {
		return res
	}

	err = s.menu.UpdateProduct(ctx, product.Product{
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

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
