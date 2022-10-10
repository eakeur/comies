package rest

import (
	"comies/app/core/ingredient"
	"comies/app/core/item"
	"comies/app/core/movement"
	"comies/app/core/order"
	"comies/app/core/product"
	"net/http"
)

var (
	failures = ErrorBinding{
		product.ErrNotFound: ResponseWithError(http.StatusNotFound, Error{
			Code:    "PRODUCT_NOT_FOUND",
			Message: "Ops! This product does not exist or could not be found",
		}),
		product.ErrCodeAlreadyExists: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "PRODUCT_CODE_ALREADY_EXISTS",
			Message: "Ops! The code assigned to this product seems to belong to another product already",
		}),
		product.ErrStockNegative: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "PRODUCT_STOCK_EMPTY",
			Message: "Ops! This product has not enough available stock",
		}),
		product.ErrStockAlreadyFull: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "PRODUCT_STOCK_FULL",
			Message: "Ops! This product's stock is already full",
		}),
		product.ErrMinimumSaleQuantity: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "PRODUCT_ZERO_SALE_QUANTITY",
			Message: "Ops! The minimum quantity of a product sale should be greater than 0",
		}),
		product.ErrInvalidCode: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "PRODUCT_INVALID_CODE",
			Message: "Ops! The product code must be longer than 2 and shorter than 12 characters",
		}),
		product.ErrInvalidType: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "PRODUCT_INVALID_TYPE",
			Message: "Ops! The product type must be a valid one",
		}),
		product.ErrInvalidName: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "PRODUCT_INVALID_NAME",
			Message: "Ops! The product name must be longer than 2 and shorter than 60 characters",
		}),
		product.ErrInvalidPrice: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "PRODUCT_ZERO_PRICE",
			Message: "Ops! The cost/sale price of a product should be greater than 0",
		}),

		ingredient.ErrInvalidIngredientID: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "INGREDIENT_INVALID_INGREDIENT_ID",
			Message: "Ops! The ingredient ID provided is invalid or does not exist",
		}),
		ingredient.ErrInvalidProductID: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "INGREDIENT_INVALID_PRODUCT_ID",
			Message: "Ops! The product ID provided is invalid or does not exist",
		}),
		ingredient.ErrInvalidQuantity: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "INGREDIENT_ZERO_QUANTITY",
			Message: "Ops! The quantity of this ingredient should be greater than zero",
		}),
		ingredient.ErrInvalidCompositeType: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "INGREDIENT_INVALID_PRODUCT_TYPE",
			Message: "a product can only have ingredients if it is of composite type",
		}),
		ingredient.ErrInvalidIngredientType: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "INGREDIENT_INVALID_INGREDIENT_TYPE",
			Message: "an output product can not compose another product",
		}),

		movement.ErrInvalidPeriod: ResponseWithError(http.StatusBadRequest, Error{
			Code:    "MOVEMENT_INVALID_PERIOD_FILTER",
			Message: "Ops! The date period informed as a filter is invalid",
		}),
		movement.ErrInvalidProductType: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "MOVEMENT_INVALID_PRODUCT_TYPE",
			Message: "output movements can not be assigned to input or composite products",
		}),
		movement.ErrInvalidType: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "MOVEMENT_INVALID_TYPE",
			Message: "the movement type must be a valid one",
		}),

		item.ErrInvalidQuantity: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "ITEM_INVALID_QUANTITY",
			Message: "Ops! This item's quantity should be greater than 0 to be ordered",
		}),
		item.ErrInvalidStatus: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "ITEM_INVALID_STATUS",
			Message: "Ops! The item status must be valid",
		}),

		order.ErrAlreadyOrdered: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "ORDER_ALREADY_ORDERED",
			Message: "Ops! This order is already in process and can not be re-ordered",
		}),
		order.ErrAlreadyPreparing: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "ORDER_ALREADY_PREPARING",
			Message: "Ops! This order is already in process and can not be canceled",
		}),
		order.ErrInvalidNumberOfItems: ResponseWithError(http.StatusPreconditionFailed, Error{
			Code:    "ORDER_MUST_HAVE_ITEMS",
			Message: "Ops! This order has no items to be ordered",
		}),
		order.ErrInvalidStatus: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "ORDER_INVALID_STATUS",
			Message: "Ops! This order status is invalid",
		}),
		order.ErrInvalidDeliveryMode: ResponseWithError(http.StatusUnprocessableEntity, Error{
			Code:    "ORDER_INVALID_DELIVERY_MODE",
			Message: "Ops! This order delivery mode is invalid",
		}),
	}.Default(ResponseWithError(http.StatusInternalServerError, Error{
		Code:    "ERR_INTERNAL_SERVER_ERROR",
		Message: "Ops! An unexpected error happened here. Please try again later",
	}))
)
