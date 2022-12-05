package ordering

import (
	"comies/core/ordering/status"
	"comies/core/types"
	"context"
)

func (w jobs) SetOrderStatus(ctx context.Context, id types.ID, st status.Status) error {
	_, err := st.Validate()
	if err != nil {
		return err
	}

	err = w.statuses.Update(ctx, st)
	if err != nil {
		return err
	}

	return nil
}
