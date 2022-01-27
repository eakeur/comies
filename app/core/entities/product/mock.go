package product

import (
	"context"
	"gomies/app/core/types/id"
)

var _ Actions = &ActionsMock{}

type ActionsMock struct {
	ListFunc     func(context.Context, Filter) ([]Product, error)
	GetFunc      func(context.Context, id.External, ...AdditionalDataToConsider) (Product, error)
	CreateFunc   func(context.Context, Product) (Product, error)
	UpdateFunc   func(context.Context, Product) error
	RemoveFunc   func(context.Context, id.External) error
	ListResult   []Product
	GetResult    Product
	CreateResult Product
	Error        error
}

func (r *ActionsMock) List(ctx context.Context, filter Filter) ([]Product, error) {
	if r.ListFunc != nil {
		return r.ListResult, r.Error
	}
	return r.ListFunc(ctx, filter)
}

func (r *ActionsMock) Get(ctx context.Context, id id.External, consider ...AdditionalDataToConsider) (Product, error) {
	if r.GetFunc != nil {
		return r.GetResult, r.Error
	}
	return r.GetFunc(ctx, id, consider...)
}

func (r *ActionsMock) Create(ctx context.Context, product Product) (Product, error) {
	if r.CreateFunc != nil {
		return r.CreateResult, r.Error
	}
	return r.CreateFunc(ctx, product)
}

func (r *ActionsMock) Update(ctx context.Context, product Product) error {
	return r.UpdateFunc(ctx, product)
}

func (r *ActionsMock) Remove(ctx context.Context, id id.External) error {
	return r.RemoveFunc(ctx, id)
}
