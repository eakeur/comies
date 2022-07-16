package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) GetProductByID(ctx context.Context, in *menu.GetProductByIDRequest) (*menu.GetProductByIDResponse, error) {
	prd, err := s.menu.GetProductByID(ctx, types.ID(in.Id))
	if err != nil {
		return nil, throw.Error(err)
	}

	return &menu.GetProductByIDResponse{
		Id:      int64(prd.ID),
		Code:    prd.Code,
		Name:    prd.Name,
		Type:    ExternalProductType(prd.Type),
		Cost:    int64(prd.CostPrice),
		Price:   int64(prd.SalePrice),
		Minimum: int64(prd.MinimumSale),
	}, nil
}
