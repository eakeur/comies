package status

import "comies/app/core/types"

func ValidateStatus(s types.Status) error {
	switch s {
	case InTheCartStatus,
		PendingStatus,
		PreparingStatus,
		WaitingTakeoutStatus,
		WaitingDeliveryStatus,
		DeliveringStatus,
		FinishedStatus,
		CanceledStatus:
		return nil
	default:
		return ErrInvalidStatus
	}
}