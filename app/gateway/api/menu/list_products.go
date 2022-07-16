package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"context"
)

func (s service) ListProducts(ctx context.Context, in *menu.ListProductsRequest) (*menu.ListProductsResponse, error) {
	prd, err := s.menu.ListProducts(ctx, product.Filter{
		Code: in.Code,
		Name: in.Name,
		Type: InternalProductType(in.Type),
	})
	if err != nil {
		return nil, throw.Error(err)
	}

	var products []*menu.ProductsListItem
	for _, p := range prd {
		products = append(products, &menu.ProductsListItem{
			Id:   int64(p.ID),
			Code: p.Code,
			Name: p.Name,
			Type: ExternalProductType(p.Type),
		})
	}

	return &menu.ListProductsResponse{
		Products: products,
	}, nil
}
