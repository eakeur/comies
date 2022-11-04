package ordering

import (
	"comies/app/core/ordering/status"
	"comies/app/core/types"
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
