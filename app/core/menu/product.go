package menu

import (
	"comies/app/core/types"
)

type Product struct {
	ID              types.ID       `json:"id"`
	Code            string         `json:"code"`
	Name            string         `json:"name"`
	Type            Type           `json:"type"`
	Balance         types.Quantity `json:"balance"`
	CostPrice       types.Currency `json:"cost_price"`
	SalePrice       types.Currency `json:"sale_price"`
	SaleUnit        types.UnitType `json:"sale_unit"`
	MinimumSale     types.Quantity `json:"minimum_sale"`
	MaximumQuantity types.Quantity `json:"maximum_quantity"`
	MinimumQuantity types.Quantity `json:"minimum_quantity"`
	Location        string         `json:"location"`
}

func ValidateProduct(p Product) error {

	if len(p.Code) < 2 || len(p.Code) > 12 {
		return ErrInvalidCode
	}

	if len(p.Name) < 2 || len(p.Name) > 60 {
		return ErrInvalidName
	}

	if err := ValidateProductType(p.Type); err != nil {
		return err
	}

	if (IsOutput(p.Type) && p.SalePrice <= 0) || p.CostPrice <= 0 {
		return ErrInvalidPrice
	}

	if IsOutput(p.Type) && p.MinimumSale <= types.QuantityMinimum {
		return ErrMinimumSaleQuantity
	}

	return nil
}

func ValidateProductType(t Type) error {
	switch t {
	case InputProductType, OutputProductType, InputCompositeProductType, OutputCompositeProductType:
		return nil
	default:
		return ErrInvalidType
	}
}

func IsComposite(t Type) bool {
	return t == InputCompositeProductType || t == OutputCompositeProductType
}

func IsInput(t Type) bool {
	return t == InputCompositeProductType || t == InputProductType
}

func IsOutput(t Type) bool {
	return t == OutputProductType || t == OutputCompositeProductType
}
