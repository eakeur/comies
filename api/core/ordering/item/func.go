package item

import "comies/core/types"

func ValidateItemStatus(s types.Status) error {
	switch s {
	case PreparingItemStatus, DoneItemStatus, FailedItemStatus:
		return nil
	default:
		return ErrInvalidStatus
	}
}
