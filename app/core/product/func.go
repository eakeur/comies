package product

import (
	"comies/app/core/types"
)

func (p Product) WithID(id types.ID) Product {
	p.ID = id
	return p
}

func (p Product) Validate() (Product, error) {

	if len(p.Code) < 2 || len(p.Code) > 12 {
		return p, ErrInvalidCode
	}

	if len(p.Name) < 2 || len(p.Name) > 60 {
		return p, ErrInvalidName
	}

	if err := ValidateType(p.Type); err != nil {
		return p, err
	}

	if (p.IsOutput() && p.SalePrice <= 0) || p.CostPrice <= 0 {
		return p, ErrInvalidPrice
	}

	if p.IsOutput() && p.MinimumSale <= types.QuantityMinimum {
		return p, ErrMinimumSaleQuantity
	}

	return p, nil
}

func ValidateType(t types.Type) error {
	switch t {
	case InputType, OutputType, InputCompositeType, OutputCompositeType:
		return nil
	default:
		return ErrInvalidType
	}
}

func (p Product) IsComposite() bool {
	return p.Type == InputCompositeType || p.Type == OutputCompositeType
}

func (p Product) IsInput() bool {
	return p.Type == InputCompositeType || p.Type == InputType
}

func (p Product) IsOutput() bool {
	return p.Type == OutputType || p.Type == OutputCompositeType
}
