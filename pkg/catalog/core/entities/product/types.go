package product

import (
	"gomies/pkg/sdk/types"
	"gomies/pkg/stocking/core/entities/stock"
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

	StockAdditionResult struct {
		RemainingStock types.Quantity
		Movement       stock.Movement
	}
)
