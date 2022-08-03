package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s Service) GetProductByKey(ctx context.Context, r *http.Request) handler.Response {
	var (
		prd product.Product
		err error
		key = chi.URLParam(r, "product_key")
	)

	// Checks if the consumer is searching by code
	if flag := r.URL.Query().Get("code"); flag == "true" {
		prd, err = s.menu.GetProductByCode(ctx, key)
	} else {
		id, e, res := handler.ConvertToID(key)
		if e != nil {
			return res
		}
		prd, err = s.menu.GetProductByID(ctx, id)
	}

	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusOK, Product{
		ID:              prd.ID.String(),
		Code:            prd.Code,
		Name:            prd.Name,
		Type:            prd.Type,
		CostPrice:       prd.CostPrice,
		SalePrice:       prd.SalePrice,
		SaleUnit:        prd.SaleUnit,
		MinimumSale:     prd.MinimumSale,
		MaximumQuantity: prd.MaximumQuantity,
		MinimumQuantity: prd.MinimumQuantity,
		Location:        prd.Location,
	})
}
