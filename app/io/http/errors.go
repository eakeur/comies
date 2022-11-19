package http

import (
	"comies/app/core/menu/ingredient"
	"comies/app/core/menu/movement"
	"comies/app/core/menu/product"
	"comies/app/core/ordering/item"
	"comies/app/core/ordering/order"
	"comies/app/core/ordering/status"
	"comies/app/io/http/send"
	"net/http"
)

func init() {
	send.RegisterDomainErrorBindings(map[error]send.Response{
		product.ErrNotFound: send.Data(http.StatusNotFound, send.ResponseError{
			Code:    "PRODUCT_NOT_FOUND",
			Message: "Ops! This product does not exist or could not be found",
		}),
		product.ErrCodeAlreadyExists: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "PRODUCT_CODE_ALREADY_EXISTS",
			Message: "Ops! The code assigned to this product seems to belong to another product already",
		}),
		product.ErrMinimumSaleQuantity: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_ZERO_SALE_QUANTITY",
			Message: "Ops! The minimum quantity of a product sale should be greater than 0",
		}),
		product.ErrInvalidCode: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_INVALID_CODE",
			Message: "Ops! The product code must be longer than 2 and shorter than 12 characters",
		}),
		product.ErrInvalidType: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_INVALID_TYPE",
			Message: "Ops! The product type must be a valid one",
		}),
		product.ErrInvalidName: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_INVALID_NAME",
			Message: "Ops! The product name must be longer than 2 and shorter than 60 characters",
		}),
		product.ErrInvalidPrice: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_ZERO_PRICE",
			Message: "Ops! The cost/sale price of a product should be greater than 0",
		}),

		ingredient.ErrInvalidComponentID: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "INGREDIENT_INVALID_INGREDIENT_ID",
			Message: "Ops! The ingredient ID provided is invalid or does not exist",
		}),
		ingredient.ErrInvalidCompositeID: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "INGREDIENT_INVALID_PRODUCT_ID",
			Message: "Ops! The product ID provided is invalid or does not exist",
		}),
		ingredient.ErrInvalidQuantity: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "INGREDIENT_ZERO_QUANTITY",
			Message: "Ops! The quantity of this ingredient should be greater than zero",
		}),

		movement.ErrInvalidPeriod: send.Data(http.StatusBadRequest, send.ResponseError{
			Code:    "MOVEMENT_INVALID_PERIOD_FILTER",
			Message: "Ops! The date period informed as a filter is invalid",
		}),
		movement.ErrInvalidProductType: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "MOVEMENT_INVALID_PRODUCT_TYPE",
			Message: "output movements can not be assigned to input or composite products",
		}),
		product.ErrInvalidType: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "MOVEMENT_INVALID_TYPE",
			Message: "the movement type must be a valid one",
		}),

		item.ErrInvalidQuantity: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "ITEM_INVALID_QUANTITY",
			Message: "Ops! This item's quantity should be greater than 0 to be ordered",
		}),
		item.ErrInvalidStatus: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "ITEM_INVALID_STATUS",
			Message: "Ops! The item status must be valid",
		}),

		order.ErrAlreadyOrdered: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "ORDER_ALREADY_ORDERED",
			Message: "Ops! This order is already in process and can not be re-ordered",
		}),
		order.ErrAlreadyPrepared: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "ORDER_ALREADY_PREPARED",
			Message: "Ops! This order is already in process and can not be canceled",
		}),
		order.ErrInvalidNumberOfItems: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "ORDER_MUST_HAVE_ITEMS",
			Message: "Ops! This order has no items to be ordered",
		}),
		status.ErrInvalidStatus: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "ORDER_INVALID_STATUS",
			Message: "Ops! This order status is invalid",
		}),
		order.ErrInvalidDeliveryType: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "ORDER_INVALID_DELIVERY_MODE",
			Message: "Ops! This order delivery mode is invalid",
		}),
	})
}
