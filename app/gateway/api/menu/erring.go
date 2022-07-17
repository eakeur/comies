package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	failures = errors.ErrorBinding{
		product.ErrNotFound:          status.New(codes.NotFound, "Ops! This product does not exist or could not be found").Err(),
		product.ErrCodeAlreadyExists: status.New(codes.AlreadyExists, "Ops! The code assigned to this product seems to belong to another product already").Err(),
	}.Default(status.New(codes.Unknown, "Ops! An unexpected error happened here. Please try again later").Err())
)
