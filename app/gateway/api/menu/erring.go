package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/gateway/api/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	failures = errors.ErrorBinding{
		product.ErrNotFound:            status.New(codes.NotFound, "Ops! This product does not exist or could not be found").Err(),
		product.ErrCodeAlreadyExists:   status.New(codes.AlreadyExists, "Ops! The code assigned to this product seems to belong to another product already").Err(),
		product.ErrStockNegative:       status.New(codes.FailedPrecondition, "Ops! This product has not enough available stock").Err(),
		product.ErrStockAlreadyFull:    status.New(codes.FailedPrecondition, "Ops! This product's stock is already full").Err(),
		product.ErrInvalidSalePrice:    status.New(codes.InvalidArgument, "Ops! This product's sale price should be greater than 0").Err(),
		product.ErrInvalidSaleQuantity: status.New(codes.InvalidArgument, "Ops! The minimum quantity of a product sale should be greater than o").Err(),

		ingredient.ErrInvalidIngredientID:   status.New(codes.InvalidArgument, "Ops! The ingredient ID provided is invalid or does not exist").Err(),
		ingredient.ErrInvalidProductID:      status.New(codes.InvalidArgument, "Ops! The product ID provided is invalid or does not exist").Err(),
		ingredient.ErrInvalidQuantity:       status.New(codes.InvalidArgument, "Ops! The quantity of this ingredient should be greater than zero").Err(),
		ingredient.ErrInvalidCompositeType:  status.New(codes.FailedPrecondition, "a product can only have ingredients if it is of composite type").Err(),
		ingredient.ErrInvalidIngredientType: status.New(codes.FailedPrecondition, "an output product can not compose another product").Err(),

		movement.ErrInvalidPeriod:      status.New(codes.InvalidArgument, "Ops! The date period informed as a filter is invalid").Err(),
		movement.ErrInvalidProductType: status.New(codes.FailedPrecondition, "output movements can not be assigned to input or composite products").Err(),
	}.Default(status.New(codes.Unknown, "Ops! An unexpected error happened here. Please try again later").Err())
)
