package menu

import (
	"comies/app/core/id"
	"comies/app/core/product"
	"comies/app/core/types"
	"comies/app/core/workflows/menu"
	"comies/app/handler/rest"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GetProductByKey fetches a product by its ID or code.
//
// @Summary     Fetches a product
// @Description Fetches a product by one of itd unique keys (id or code).
// @Tags        Product
// @Param       product_key path     string false "The product ID"
// @Param       code        query    bool   false "Toggles if the API should search by code"
// @Success     200         {object} rest.Response{data=GetProductByKeyResponse{}}
// @Failure     404         {object} rest.Response{error=rest.Error{}} "PRODUCT_NOT_FOUND"
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id} [GET]
func GetProductByKey(ctx context.Context, r *http.Request) rest.Response {
	var (
		prd product.Product
		err error
		key = chi.URLParam(r, "product_key")
	)

	// Checks if the consumer is searching by code
	if flag := r.URL.Query().Get("code"); flag == "true" {
		prd, err = menu.GetProductByCode(ctx, key)
	} else {
		var id id.ID
		id, err = rest.ConvertToID(key)
		if err != nil {
			return rest.IDParsingErrorResponse(err)
		}
		prd, err = menu.GetProductByID(ctx, id)
	}

	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusOK, NewGetProductByKeyResponse(prd))
}

type (
	GetProductByKeyResponse struct {
		// ID is the unique identifier of this product
		ID string `json:"id"`
		// Code represents how the store's crew call this product internally
		Code string `json:"code"`
		// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
		// shown in fiscal documents
		Name string `json:"name"`
		// Type is the type of the product
		Type product.Type `json:"type"`
		// CostPrice is how much the store pays to make or store this product
		CostPrice types.Currency `json:"cost_price"`
		// Price is how much the customer pays for this product
		SalePrice types.Currency `json:"sale_price"`
		// Unit is the measure type that this product is sold
		SaleUnit types.UnitType `json:"sale_unit"`
		// MinimumSale is the lowest number of unities of this product that can be sold
		MinimumSale types.Quantity `json:"minimum_sale"`
		// MaximumQuantity is how many unities of this resource the stock can support
		MaximumQuantity types.Quantity `json:"maximum_quantity"`
		// MinimumQuantity is the lowest quantity of this resource the stock can have
		MinimumQuantity types.Quantity `json:"minimum_quantity"`
		// Location is a brief description of where this stock is located
		Location string `json:"location"`
		// Balance is the stock balance of this product
		Balance types.Quantity `json:"balance"`
	}
)

func NewGetProductByKeyResponse(prd product.Product) GetProductByKeyResponse {
	return GetProductByKeyResponse{
		ID:              prd.ID.String(),
		Code:            prd.Code,
		Name:            prd.Name,
		Type:            prd.Type,
		CostPrice:       prd.CostPrice,
		SalePrice:       prd.SalePrice,
		SaleUnit:        prd.SaleUnit,
		MinimumSale:     prd.MinimumSale,
		MaximumQuantity: prd.MaximumQuantity,
		MinimumQuantity: prd.MinimumQuantity,
		Location:        prd.Location,
		Balance:         prd.Balance,
	}
}
