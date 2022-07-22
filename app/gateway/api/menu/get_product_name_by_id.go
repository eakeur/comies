package menu

import (
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) GetProductNameByID(ctx context.Context, params handler.RouteParams) response.Response {
	id, err, res := convertToID(params["product_id"])
	if err != nil {
		return res
	}

	name, err := s.menu.GetProductNameByID(ctx, id)
	if err != nil {
		return response.Response{}
	}
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusOK, ProductNameResult{Name: name})
}
