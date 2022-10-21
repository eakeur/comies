package menu

import (
	"comies/app/core/product"
)

func SaveProduct(id IDGenerator, write ProductWriter) ProductSaver {
	return func(p product.Product) (product.Product, error) {
		save, err := p.WithID(id()).Validate()
		if err != nil {
			return product.Product{}, err
		}

		return save, write(save)
	}
}
