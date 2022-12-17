package ingredient

import (
	"comies/core/types"
	"comies/data/conn"
	"context"
)

func (a actions) RemoveByProductID(ctx context.Context, productID types.ID) error {
	const script = `delete from ingredients where product_id = $1`

	_, err := conn.ExecFromContext(ctx, script, productID)
	// TODO: validate if lines affected are different from 1
	return err
}
