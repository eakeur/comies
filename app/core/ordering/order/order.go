package order

import (
	"comies/app/core/types"
	"time"
)

type CountByStatus = map[types.Status]types.Quantity

type Order struct {
	ID              types.ID   `json:"id"`
	PlacedAt        time.Time  `json:"placed_at"`
	DeliveryType    types.Type `json:"delivery_mode"`
	Observations    string     `json:"observations"`
	CustomerName    string     `json:"customer_name"`
	CustomerAddress string     `json:"customer_address"`
	CustomerPhone   string     `json:"customer_phone"`
}

func (o Order) Validate() error {
	if o.PlacedAt.IsZero() || o.PlacedAt.After(time.Now().UTC()) {
		return ErrInvalidPlacementDate
	}

	if err := ValidateDeliveryType(o.DeliveryType); err != nil {
		return err
	}

	if len(o.CustomerName) <= 0 {
		return ErrInvalidCustomerName
	}

	if len(o.CustomerPhone) <= 0 {
		return ErrInvalidCustomerPhone
	}

	if o.DeliveryType == DeliverDeliveryType {
		if len(o.CustomerAddress) <= 0 {
			return ErrInvalidCustomerAddress
		}
	}

	return nil
}
