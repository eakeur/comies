package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
	"net/url"
)

func (s Service) GetProductByKey(ctx context.Context, params handler.RouteParams, query url.Values) response.Response {
	var (
		prd product.Product
		err error
		key = params["product_key"]
	)

	// Checks if the consumer is searching by code
	if flag := query.Get("code"); flag == "true" {
		prd, err = s.menu.GetProductByCode(ctx, key)
	} else {
		id, e, res := convertToID(key)
		if e != nil {
			return res
		}
		prd, err = s.menu.GetProductByID(ctx, id)
	}

	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusOK, Product{
		ID:              prd.ID.String(),
		Code:            prd.Code,
		Name:            prd.Name,
		Type:            prd.Type,
		CostPrice:       prd.CostPrice,
		SalePrice:       prd.SalePrice,
		MinimumSale:     prd.MinimumSale,
		MaximumQuantity: prd.MaximumQuantity,
		MinimumQuantity: prd.MinimumQuantity,
	})
}
