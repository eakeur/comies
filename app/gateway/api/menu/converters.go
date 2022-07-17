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
	}
)

func InternalProduct(p ProductInfoGetter) product.Product {
	prd := product.Product{
		Code: p.GetCode(),
		Name: p.GetName(),
		Type: InternalProductType(p.GetType()),
		Sale: product.Sale{
			CostPrice:   types.Currency(p.GetCost()),
			SalePrice:   types.Currency(p.GetPrice()),
			MinimumSale: types.Quantity(p.GetMinimum()),
		},
	}

	return prd
}

func ExternalProductType(p product.Type) menu.ProductType {
	switch p {
	case product.InputType:
		return menu.ProductType_INPUT
	case product.OutputType:
		return menu.ProductType_INPUT
	default:
		return menu.ProductType_OUTPUT
	}
}

func InternalProductType(p menu.ProductType) product.Type {
	return product.Type(p.Descriptor().Name())
}
