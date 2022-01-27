package category

import (
	"context"
	"gomies/app/core/types/id"
)

var _ Actions = ActionsMock{}

type ActionsMock struct {
	GetResult Category
	ListResult []Category
	CreateResult Category
	Error error
}

func (a ActionsMock) Get(_ context.Context, _ id.External) (Category, error) {
	return a.GetResult, a.Error
}

func (a ActionsMock) List(_ context.Context, _ Filter) ([]Category, error) {
	return a.ListResult, a.Error
}

func (a ActionsMock) Create(_ context.Context, _ Category) (Category, error) {
	return a.CreateResult, a.Error
}

func (a ActionsMock) Remove(_ context.Context, _ id.External) error {
	return a.Error
}
