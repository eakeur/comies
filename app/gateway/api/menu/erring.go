package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/gateway/api/response"
	"net/http"
)

var (
	failures = response.ErrorBinding{
		product.ErrNotFound: response.WithError(http.StatusNotFound, response.Error{
			Code:    "PRODUCT_NOT_FOUND",
			Message: "Ops! This product does not exist or could not be found",
		}),

		product.ErrCodeAlreadyExists: response.WithError(http.StatusPreconditionFailed, response.Error{
			Code:    "PRODUCT_CODE_ALREADY_EXISTS",
			Message: "Ops! The code assigned to this product seems to belong to another product already",
		}),
		product.ErrStockNegative: response.WithError(http.StatusPreconditionFailed, response.Error{
			Code:    "PRODUCT_STOCK_EMPTY",
			Message: "Ops! This product has not enough available stock",
		}),
		product.ErrStockAlreadyFull: response.WithError(http.StatusPreconditionFailed, response.Error{
			Code:    "PRODUCT_STOCK_FULL",
			Message: "Ops! This product's stock is already full",
		}),
		product.ErrInvalidSalePrice: response.WithError(http.StatusUnprocessableEntity, response.Error{
			Code:    "PRODUCT_ZERO_SALE_PRICE",
			Message: "Ops! This product's sale price should be greater than 0",
		}),
		product.ErrInvalidSaleQuantity: response.WithError(http.StatusUnprocessableEntity, response.Error{
			Code:    "PRODUCT_ZERO_SALE_QUANTITY",
			Message: "Ops! The minimum quantity of a product sale should be greater than 0",
		}),

		ingredient.ErrInvalidIngredientID: response.WithError(http.StatusPreconditionFailed, response.Error{
			Code:    "INGREDIENT_INVALID_INGREDIENT_ID",
			Message: "Ops! The ingredient ID provided is invalid or does not exist",
		}),
		ingredient.ErrInvalidProductID: response.WithError(http.StatusPreconditionFailed, response.Error{
			Code:    "INGREDIENT_INVALID_PRODUCT_ID",
			Message: "Ops! The product ID provided is invalid or does not exist",
		}),
		ingredient.ErrInvalidQuantity: response.WithError(http.StatusUnprocessableEntity, response.Error{
			Code:    "INGREDIENT_ZERO_QUANTITY",
			Message: "Ops! The quantity of this ingredient should be greater than zero",
		}),
		ingredient.ErrInvalidCompositeType: response.WithError(http.StatusPreconditionFailed, response.Error{
			Code:    "INGREDIENT_INVALID_PRODUCT_TYPE",
			Message: "a product can only have ingredients if it is of composite type",
		}),
		ingredient.ErrInvalidIngredientType: response.WithError(http.StatusPreconditionFailed, response.Error{
			Code:    "INGREDIENT_INVALID_INGREDIENT_TYPE",
			Message: "an output product can not compose another product",
		}),

		movement.ErrInvalidPeriod: response.WithError(http.StatusBadRequest, response.Error{
			Code:    "MOVEMENT_INVALID_PERIOD_FILTER",
			Message: "Ops! The date period informed as a filter is invalid",
		}),
		movement.ErrInvalidProductType: response.WithError(http.StatusPreconditionFailed, response.Error{
			Code:    "MOVEMENT_INVALID_PRODUCT_TYPE",
			Message: "output movements can not be assigned to input or composite products",
		}),
	}.Default(response.WithError(http.StatusInternalServerError, response.Error{
		Code:    "ERR_INTERNAL_SERVER_ERROR",
		Message: "Ops! An unexpected error happened here. Please try again later",
	}))
)
