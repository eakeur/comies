package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"context"
)

func (s service) GetProductByCode(ctx context.Context, request *menu.GetProductByCodeRequest) (*menu.GetProductByCodeResponse, error) {
	prd, err := s.menu.GetProductByCode(ctx, request.Code)
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.GetProductByCodeResponse{
		Product: &menu.Product{
			Id:      int64(prd.ID),
			Code:    prd.Code,
			Name:    prd.Name,
			Type:    menu.ProductType(prd.Type),
			Cost:    int64(prd.CostPrice),
			Price:   int64(prd.SalePrice),
			Minimum: int64(prd.MinimumSale),
		},
	}, nil
}
