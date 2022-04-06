package stock

import (
	"context"
	"gomies/app/core/entities/stocking/stock"
	"gomies/pkg/sdk/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_ClosePeriod(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

	type (
		args struct {
			filter stock.Filter
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
			name: "should close period successfully",
			args: args{
				filter: stock.Filter{
					ResourceID: fakeID,
					FinalDate:  time.Now(),
				},
			},
			opts: opts{
				stocks: &stock.ActionsMock{
					ArchiveMovementsFunc: func(ctx context.Context, filter stock.Filter) error {
						return nil
					},
				},
			},
		},
		{
			name: "should fail because resourceID is invalid",
			args: args{
				filter: stock.Filter{
					ResourceID: "",
					FinalDate:  time.Now(),
				},
			},
			wantErr: stock.ErrMissingResourceID,
			opts: opts{
				stocks: &stock.ActionsMock{
					ArchiveMovementsFunc: func(ctx context.Context, filter stock.Filter) error {
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
			err := wf.ClosePeriod(ctx, c.args.filter)
			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
