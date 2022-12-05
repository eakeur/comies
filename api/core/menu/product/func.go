package product

import "comies/core/types"

func ValidateType(t types.Type) error {
	switch t {
	case InputType, OutputType, InputCompositeType, OutputCompositeType:
		return nil
	default:
		return ErrInvalidType
	}
}
