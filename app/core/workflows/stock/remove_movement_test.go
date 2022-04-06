package stock

import (
	"context"
	"gomies/app/core/entities/stocking/stock"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_RemoveMovement(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

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
				resourceID: fakeID,
				movementID: fakeID,
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
				resourceID: fakeID,
				movementID: "",
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
