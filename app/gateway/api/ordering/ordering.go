package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/response"
)

var failures = response.ErrorBinding{}

type Service struct {
	ordering ordering.Workflow
}

func NewService(ordering ordering.Workflow) *Service {
	return &Service{
		ordering: ordering,
	}
}
