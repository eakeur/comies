package menu

import (
	"comies/app/core/entities/movement"
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
		Type: InternalProductType(p.GetType()),
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

func ExternalMovementType(p movement.Type) (t menu.MovementType) {
	switch p {
	case movement.InputMovement:
		return menu.MovementType_INPUT_MOVEMENT_TYPE
	case movement.OutputMovement:
		return menu.MovementType_OUTPUT_MOVEMENT_TYPE
	case movement.ReservedMovement:
		return menu.MovementType_RESERVED_MOVEMENT_TYPE
	}
	return t
}

func InternalMovementType(p menu.MovementType) (t movement.Type) {
	switch p {
	case menu.MovementType_INPUT_MOVEMENT_TYPE:
		return movement.InputMovement
	case menu.MovementType_OUTPUT_MOVEMENT_TYPE:
		return movement.OutputMovement
	case menu.MovementType_RESERVED_MOVEMENT_TYPE:
		return movement.ReservedMovement
	}
	return t
}
