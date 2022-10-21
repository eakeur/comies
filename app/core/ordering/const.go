package ordering

type (
	Type   int
	Status int
)

const (
	NoType   Type   = 0
	NoStatus Status = 0
)

const (
	InTheCartOrderStatus       Status = 10
	PendingOrderStatus         Status = 20
	PreparingOrderStatus       Status = 30
	WaitingTakeoutOrderStatus  Status = 40
	WaitingDeliveryOrderStatus Status = 50
	DeliveringOrderStatus      Status = 60
	FinishedOrderStatus        Status = 70
	CanceledOrderStatus        Status = 80
)

const (
	PreparingItemStatus Status = 10
	DoneItemStatus      Status = 20
	FailedItemStatus    Status = 30
)

const (
	TakeoutDeliveryType Type = 10
	DeliverDeliveryType Type = 20
)
