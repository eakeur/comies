package product

import (
	"gomies/pkg/sdk/types"
)

type (
	ApproveSaleRequest struct {
		Key
		Quantity types.Quantity
		Price    types.Currency
		Ignore   bool
		Replace  Key
	}

	ApproveSaleResponse struct {
		RemainingStock types.Quantity
		Price          types.Quantity
	}
)
