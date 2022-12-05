package product

import "comies/core/types"

type Product struct {
	ID              types.ID       `json:"id"`
	Code            string         `json:"code"`
	Name            string         `json:"name"`
	Type            types.Type     `json:"type"`
	CostPrice       types.Currency `json:"cost_price"`
	SaleUnit        types.UnitType `json:"sale_unit"`
	MinimumSale     types.Quantity `json:"minimum_sale"`
	MaximumQuantity types.Quantity `json:"maximum_quantity"`
	MinimumQuantity types.Quantity `json:"minimum_quantity"`
	Location        string         `json:"location"`
}

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

	if p.CostPrice <= 0 {
		return p, ErrInvalidPrice
	}

	if p.IsOutput() && p.MinimumSale <= types.QuantityMinimum {
		return p, ErrMinimumSaleQuantity
	}

	return p, nil
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
