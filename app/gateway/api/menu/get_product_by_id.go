package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) GetProductById(ctx context.Context, in *menu.GetProductByIdRequest) (*menu.GetProductByIdResponse, error) {
	prd, err := s.menu.GetProductByID(ctx, types.ID(in.Id))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.GetProductByIdResponse{
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
