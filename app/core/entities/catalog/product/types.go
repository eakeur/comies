package product

import (
	"gomies/pkg/sdk/types"
)

type (
	ApproveSaleRequest struct {
		Key
		Quantity types.Quantity
		Price    types.Currency
	}

	ApproveSaleResponse struct {
		RemainingStock types.Quantity
		Price          types.Quantity
	}
)
