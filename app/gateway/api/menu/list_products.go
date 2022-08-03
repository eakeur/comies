package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
	"strconv"
)

func (s Service) ListProducts(ctx context.Context, r *http.Request) handler.Response {
	query := r.URL.Query()
	filter := product.Filter{
		Code: query.Get("code"),
		Name: query.Get("name"),
		Type: 0,
	}

	runningOut := query.Get("out") == "true"

	t, err := strconv.Atoi(query.Get("type"))
	if err == nil {
		filter.Type = product.Type(t)
	}

	var prd []product.Product
	if runningOut {
		prd, err = s.menu.ListProductsRunningOut(ctx)
	} else {
		prd, err = s.menu.ListProducts(ctx, filter)
	}
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	products := make([]Product, len(prd))
	for i, p := range prd {
		products[i] = Product{
			ID:      p.ID.String(),
			Code:    p.Code,
			Name:    p.Name,
			Type:    p.Type,
			Balance: p.Balance,
		}
	}

	return handler.ResponseWithData(http.StatusOK, products)
}
