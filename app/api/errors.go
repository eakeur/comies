package api

import (
	"comies/app/api/send"
	"comies/app/core/menu"
	"comies/app/core/ordering"
	"net/http"
)

func init() {
	send.RegisterDomainErrorBindings(map[error]send.Response{
		menu.ErrNotFound: send.Data(http.StatusNotFound, send.ResponseError{
			Code:    "PRODUCT_NOT_FOUND",
			Message: "Ops! This product does not exist or could not be found",
		}),
		menu.ErrCodeAlreadyExists: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "PRODUCT_CODE_ALREADY_EXISTS",
			Message: "Ops! The code assigned to this product seems to belong to another product already",
		}),
		menu.ErrStockNegative: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "PRODUCT_STOCK_EMPTY",
			Message: "Ops! This product has not enough available stock",
		}),
		menu.ErrStockAlreadyFull: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "PRODUCT_STOCK_FULL",
			Message: "Ops! This product's stock is already full",
		}),
		menu.ErrMinimumSaleQuantity: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_ZERO_SALE_QUANTITY",
			Message: "Ops! The minimum quantity of a product sale should be greater than 0",
		}),
		menu.ErrInvalidCode: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_INVALID_CODE",
			Message: "Ops! The product code must be longer than 2 and shorter than 12 characters",
		}),
		menu.ErrInvalidType: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_INVALID_TYPE",
			Message: "Ops! The product type must be a valid one",
		}),
		menu.ErrInvalidName: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_INVALID_NAME",
			Message: "Ops! The product name must be longer than 2 and shorter than 60 characters",
		}),
		menu.ErrInvalidPrice: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "PRODUCT_ZERO_PRICE",
			Message: "Ops! The cost/sale price of a product should be greater than 0",
		}),

		menu.ErrInvalidComponentID: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "INGREDIENT_INVALID_INGREDIENT_ID",
			Message: "Ops! The ingredient ID provided is invalid or does not exist",
		}),
		menu.ErrInvalidCompositeID: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "INGREDIENT_INVALID_PRODUCT_ID",
			Message: "Ops! The product ID provided is invalid or does not exist",
		}),
		menu.ErrInvalidQuantity: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "INGREDIENT_ZERO_QUANTITY",
			Message: "Ops! The quantity of this ingredient should be greater than zero",
		}),

		menu.ErrInvalidPeriod: send.Data(http.StatusBadRequest, send.ResponseError{
			Code:    "MOVEMENT_INVALID_PERIOD_FILTER",
			Message: "Ops! The date period informed as a filter is invalid",
		}),
		menu.ErrInvalidProductType: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "MOVEMENT_INVALID_PRODUCT_TYPE",
			Message: "output movements can not be assigned to input or composite products",
		}),
		menu.ErrInvalidType: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "MOVEMENT_INVALID_TYPE",
			Message: "the movement type must be a valid one",
		}),

		ordering.ErrInvalidQuantity: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "ITEM_INVALID_QUANTITY",
			Message: "Ops! This item's quantity should be greater than 0 to be ordered",
		}),
		ordering.ErrInvalidStatus: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "ITEM_INVALID_STATUS",
			Message: "Ops! The item status must be valid",
		}),

		ordering.ErrAlreadyOrdered: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "ORDER_ALREADY_ORDERED",
			Message: "Ops! This order is already in process and can not be re-ordered",
		}),
		ordering.ErrAlreadyPreparing: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "ORDER_ALREADY_PREPARING",
			Message: "Ops! This order is already in process and can not be canceled",
		}),
		ordering.ErrInvalidNumberOfItems: send.Data(http.StatusPreconditionFailed, send.ResponseError{
			Code:    "ORDER_MUST_HAVE_ITEMS",
			Message: "Ops! This order has no items to be ordered",
		}),
		ordering.ErrInvalidStatus: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "ORDER_INVALID_STATUS",
			Message: "Ops! This order status is invalid",
		}),
		ordering.ErrInvalidDeliveryType: send.Data(http.StatusUnprocessableEntity, send.ResponseError{
			Code:    "ORDER_INVALID_DELIVERY_MODE",
			Message: "Ops! This order delivery mode is invalid",
		}),
	})
}
