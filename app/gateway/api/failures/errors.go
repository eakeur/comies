package failures

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/item"
	"comies/app/core/entities/movement"
	"comies/app/core/entities/order"
	"comies/app/core/entities/product"
	"comies/app/gateway/api/handler"
	"net/http"
)

var (
	failures = ErrorBinding{
		product.ErrNotFound: handler.ResponseWithError(http.StatusNotFound, handler.Error{
			Code:    "PRODUCT_NOT_FOUND",
			Message: "Ops! This product does not exist or could not be found",
		}),
		product.ErrCodeAlreadyExists: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "PRODUCT_CODE_ALREADY_EXISTS",
			Message: "Ops! The code assigned to this product seems to belong to another product already",
		}),
		product.ErrStockNegative: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "PRODUCT_STOCK_EMPTY",
			Message: "Ops! This product has not enough available stock",
		}),
		product.ErrStockAlreadyFull: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "PRODUCT_STOCK_FULL",
			Message: "Ops! This product's stock is already full",
		}),
		product.ErrMinimumSaleQuantity: handler.ResponseWithError(http.StatusUnprocessableEntity, handler.Error{
			Code:    "PRODUCT_ZERO_SALE_QUANTITY",
			Message: "Ops! The minimum quantity of a product sale should be greater than 0",
		}),
		product.ErrInvalidCode: handler.ResponseWithError(http.StatusUnprocessableEntity, handler.Error{
			Code:    "PRODUCT_INVALID_CODE",
			Message: "Ops! The product code must be longer than 2 and shorter than 12 characters",
		}),
		product.ErrInvalidName: handler.ResponseWithError(http.StatusUnprocessableEntity, handler.Error{
			Code:    "PRODUCT_INVALID_NAME",
			Message: "Ops! The product name must be longer than 2 and shorter than 60 characters",
		}),
		product.ErrInvalidPrice: handler.ResponseWithError(http.StatusUnprocessableEntity, handler.Error{
			Code:    "PRODUCT_ZERO_PRICE",
			Message: "Ops! The cost/sale price of a product should be greater than 0",
		}),

		ingredient.ErrInvalidIngredientID: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "INGREDIENT_INVALID_INGREDIENT_ID",
			Message: "Ops! The ingredient ID provided is invalid or does not exist",
		}),
		ingredient.ErrInvalidProductID: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "INGREDIENT_INVALID_PRODUCT_ID",
			Message: "Ops! The product ID provided is invalid or does not exist",
		}),
		ingredient.ErrInvalidQuantity: handler.ResponseWithError(http.StatusUnprocessableEntity, handler.Error{
			Code:    "INGREDIENT_ZERO_QUANTITY",
			Message: "Ops! The quantity of this ingredient should be greater than zero",
		}),
		ingredient.ErrInvalidCompositeType: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "INGREDIENT_INVALID_PRODUCT_TYPE",
			Message: "a product can only have ingredients if it is of composite type",
		}),
		ingredient.ErrInvalidIngredientType: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "INGREDIENT_INVALID_INGREDIENT_TYPE",
			Message: "an output product can not compose another product",
		}),

		movement.ErrInvalidPeriod: handler.ResponseWithError(http.StatusBadRequest, handler.Error{
			Code:    "MOVEMENT_INVALID_PERIOD_FILTER",
			Message: "Ops! The date period informed as a filter is invalid",
		}),
		movement.ErrInvalidProductType: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "MOVEMENT_INVALID_PRODUCT_TYPE",
			Message: "output movements can not be assigned to input or composite products",
		}),

		item.ErrInvalidQuantity: handler.ResponseWithError(http.StatusUnprocessableEntity, handler.Error{
			Code:    "ITEM_INVALID_QUANTITY",
			Message: "Ops! This item's quantity should be greater than 0 to be ordered",
		}),
		order.ErrAlreadyOrdered: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "ORDER_ALREADY_ORDERED",
			Message: "Ops! This order is already in process and can not be re-ordered",
		}),
		order.ErrAlreadyPreparing: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "ORDER_ALREADY_PREPARING",
			Message: "Ops! This order is already in process and can not be canceled",
		}),
		order.ErrInvalidNumberOfItems: handler.ResponseWithError(http.StatusPreconditionFailed, handler.Error{
			Code:    "ORDER_MUST_HAVE_ITEMS",
			Message: "Ops! This order has no items to be ordered",
		}),
	}.Default(handler.ResponseWithError(http.StatusInternalServerError, handler.Error{
		Code:    "ERR_INTERNAL_SERVER_ERROR",
		Message: "Ops! An unexpected error happened here. Please try again later",
	}))
)
