package status

import "comies/app/core/types"

const (
	InTheCartStatus       types.Status = 10
	PendingStatus         types.Status = 20
	PreparingStatus       types.Status = 30
	WaitingTakeoutStatus  types.Status = 40
	WaitingDeliveryStatus types.Status = 50
	DeliveringStatus      types.Status = 60
	FinishedStatus        types.Status = 70
	CanceledStatus        types.Status = 80
)

var UnfinishedStatuses = []types.Status{
	PendingStatus,
	PreparingStatus,
	WaitingTakeoutStatus,
	WaitingDeliveryStatus,
	DeliveringStatus,
}