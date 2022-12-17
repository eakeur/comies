package billing

import (
	"comies/core/billing/bill"
	"comies/core/billing/item"
	"comies/core/types"
	"context"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Jobs:WorkflowMock
type Jobs interface {
	CreateBill(ctx context.Context, b BillCreation) (BillSummary, error)
}

type jobs struct {
	bills    bill.Actions
	items    item.Actions
	createID types.CreateID
}

type Deps struct {
	Bills     bill.Actions
	Items     item.Actions
	IDCreator types.CreateID
}

var _ Jobs = jobs{}

func NewJobs(deps Deps) Jobs {
	return jobs{
		bills:    deps.Bills,
		items:    deps.Items,
		createID: deps.IDCreator,
	}
}
