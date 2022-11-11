package ingredient

import (
	"comies/app/core/types"
	"comies/app/gateway/data/postgres/conn"
	"context"
)

func (a actions) Remove(ctx context.Context, ingredientID types.ID) error {
	const script = "delete from ingredients where id = $1"

	_, err := conn.ExecFromContext(ctx, script, ingredientID)
	// TODO: validate if lines affected are different from 1
	return err
}
