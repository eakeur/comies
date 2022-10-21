package menu

import (
	"comies/app/core/product"
	"comies/app/core/types"
)

func GetProductByID(fetcher ProductFetcher) func(id types.ID) (product.Product, error) {
	return func(id types.ID) (product.Product, error) {
		if err := types.ValidateID(id); err != nil {
			return product.Product{}, err
		}

		return fetcher(id)
	}
}
