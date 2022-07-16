package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) CreateProduct(ctx context.Context, in *menu.CreateProductRequest) (*menu.CreateProductResponse, error) {
	prd, err := s.menu.CreateProduct(ctx, product.Product{
		Code: in.Code,
		Name: in.Name,
		Type: product.Type(in.Type),
		Sale: product.Sale{
			CostPrice:   types.Currency(in.Cost),
			SalePrice:   types.Currency(in.Price),
			SaleUnit:    types.UnitType(in.Unit),
			MinimumSale: types.Quantity(in.Minimum),
		},
	})
	if err != nil {
		return nil, throw.Error(err)
	}

	return &menu.CreateProductResponse{
		Id: int64(prd.ID),
	}, nil
}
