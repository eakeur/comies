package order

import (
	"comies/core/types"
	"time"
)

type Order struct {
	ID              types.ID   `json:"id"`
	PlacedAt        time.Time  `json:"placed_at"`
	DeliveryType    types.Type `json:"delivery_mode"`
	Observations    string     `json:"observations"`
	CustomerName    string     `json:"customer_name"`
	CustomerAddress string     `json:"customer_address"`
	CustomerPhone   string     `json:"customer_phone"`
}

func (o Order) WithID(id types.ID) Order {
	o.ID = id
	return o
}

func (o Order) WithPlacedAt(d time.Time) Order {
	o.PlacedAt = d.UTC()
	return o
}

func (o Order) WithDeliveryType(d types.Type) Order {
	o.DeliveryType = d
	return o
}

func (o Order) WithCustomer(name, phone, addr string) Order {
	o.CustomerName = name
	o.CustomerPhone = phone
	o.CustomerAddress = addr

	return o
}

func (o Order) WithObservations(obs string) Order {
	o.Observations = obs
	return o
}

func (o Order) Validate() (Order, error) {
	if o.PlacedAt.IsZero() || o.PlacedAt.After(time.Now().UTC()) {
		return o, ErrInvalidPlacementDate
	}

	if err := ValidateDeliveryType(o.DeliveryType); err != nil {
		return o, err
	}

	if len(o.CustomerName) <= 0 {
		return o, ErrInvalidCustomerName
	}

	if len(o.CustomerPhone) <= 0 {
		return o, ErrInvalidCustomerPhone
	}

	if o.DeliveryType == DeliverDeliveryType {
		if len(o.CustomerAddress) <= 0 {
			return o, ErrInvalidCustomerAddress
		}
	}

	return o, nil
}
