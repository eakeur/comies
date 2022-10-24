package order

import "comies/app/core/types"

func ValidateDeliveryType(t types.Type) error {
	switch t {
	case TakeoutDeliveryType, DeliverDeliveryType:
		return nil
		default:
			return ErrInvalidDeliveryType
	}
}