package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/types"
)

type (
	ProductInfoGetter interface {
		GetCode() string
		GetName() string
		GetType() menu.ProductType
		GetCost() int64
		GetPrice() int64
		GetMinimum() int64
		GetStockMaximum() int64
		GetStockMinimum() int64
		GetLocation() string
	}
)

func InternalProduct(p ProductInfoGetter) product.Product {
	prd := product.Product{
		Code: p.GetCode(),
		Name: p.GetName(),
		Type: product.Type(p.GetType().Number()),
		Sale: product.Sale{
			CostPrice:   types.Currency(p.GetCost()),
			SalePrice:   types.Currency(p.GetPrice()),
			MinimumSale: types.Quantity(p.GetMinimum()),
		},
		Stock: product.Stock{
			MaximumQuantity: types.Quantity(p.GetStockMaximum()),
			MinimumQuantity: types.Quantity(p.GetStockMinimum()),
			Location:        p.GetLocation(),
		},
	}

	return prd
}
