package product

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/core/entities/stocking/stock"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {

	// SaveProduct saves a product or updates an existing one
	//
	// Possible errors:
	//   - session.ErrNoSession: if there is no session in this context
	//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
	//   - fault.ErrAlreadyExists: if the product already exists and the "overwrite" flag was not set
	//   - ErrInvalidCode: if the code is invalid
	//   - ErrInvalidName: if the name is invalid
	//   - ErrInvalidPrice: if the sal price is 0
	SaveProduct(ctx context.Context, prd product.Product, flag ...types.WritingFlag) (product.Product, error)

	// ListProducts retrieves all products respecting a given filter
	ListProducts(ctx context.Context, productFilter product.Filter) ([]product.Product, error)

	// GetProduct retrieves a product respecting its key
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the product is not found
	GetProduct(ctx context.Context, key product.Key) (product.Product, error)

	// RemoveProduct deletes a product or deactivates it if there is some children
	RemoveProduct(ctx context.Context, key product.Key) error

	// ApproveSale checks if a product can be sold with such parameters
	//
	// Possible errors
	//   - session.ErrNoSession: if there is no session in this context
	//   - session.ErrNotAllowed: if the session owner is not allowed to perform this operation
	//   - fault.ErrNotFound: if the product does not exist
	//   - product.ErrInvalidSalePrice: if the price requested is not allowed to be used
	//   - product.ErrInvalidSaleQuantity: if the quantity requested is not allowed to be used
	//   - product.ErrNotEnoughStocked: if the quantity requested is not afforded by the stock
	ApproveSale(ctx context.Context, req product.ApproveSaleRequest) error

	AddIngredient(ctx context.Context, productKey product.Key, ingredient product.Ingredient) (product.Ingredient, error)

	RemoveIngredient(ctx context.Context, productKey product.Key, id types.UID) error
}

var _ Workflow = workflow{}

func NewWorkflow(
	products product.Actions,
	categories category.Actions,
	stocks stock.Workflow,
) Workflow {
	return workflow{
		products:   products,
		stocks:     stocks,
		categories: categories,
	}
}

type workflow struct {
	products   product.Actions
	categories category.Actions
	stocks     stock.Workflow
}
