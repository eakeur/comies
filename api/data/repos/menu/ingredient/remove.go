package ingredient

import (
	"comies/core/types"
	"comies/data/conn"
	"context"
)

func (a actions) Remove(ctx context.Context, productID types.ID, ingredientID types.ID) error {
	const script = "delete from ingredients where product_id = $1 and ingredient_id = $2"

	_, err := conn.ExecFromContext(ctx, script, productID, ingredientID)
	// TODO: validate if lines affected are different from 1
	return err
}
