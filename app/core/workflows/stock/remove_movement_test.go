package stock

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/stocking/stock"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_RemoveMovement(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			resourceID types.UID
			movementID types.UID
		}

		opts struct {
			stocks *stock.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should remove movement successfully",
			args: args{
				resourceID: idExample1,
				movementID: idExample2,
			},
			opts: opts{
				stocks: &stock.ActionsMock{
					RemoveMovementFunc: func(ctx context.Context, resourceID types.UID, movementID types.UID) error {
						return nil
					},
				},
			},
		},
		{
			name: "should fail because movementID is invalid",
			args: args{
				resourceID: idExample1,
				movementID: types.UID{},
			},
			wantErr: stock.ErrMissingResourceID,
			opts: opts{
				stocks: &stock.ActionsMock{
					RemoveMovementFunc: func(ctx context.Context, resourceID types.UID, movementID types.UID) error {
						return nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.stocks)
			err := wf.RemoveMovement(ctx, c.args.resourceID, c.args.movementID)
			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
