package bill

import (
	"comies/core/types"
	"context"
)

type Actions interface {
	Create(ctx context.Context, b Bill) error
	List(ctx context.Context, f Filter) ([]Bill, error)
	GetByID(ctx context.Context, id types.ID) (Bill, error)
	GetByReferenceID(ctx context.Context, ref types.ID) (Bill, error)
}
