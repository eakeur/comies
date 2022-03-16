package wrappers

import (
	"gomies/app/core/entities/category"
	"gomies/app/core/entities/crew"
	"gomies/app/core/entities/product"
)

type Workflows struct {
	Products   product.Workflow
	Categories category.Workflow
	Crew       crew.Workflow
}
