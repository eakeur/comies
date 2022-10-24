package movement

import "comies/app/core/types"

func ValidateType(t types.Type) error {
	switch t {
	case InputType, OutputType:
		return nil
	default:
		return ErrInvalidType
	}
}
