package stock

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/stocking/stock"
	"gomies/pkg/sdk/types"
	"testing"
	"time"
)

func TestWorkflow_ClosePeriod(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

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
					ResourceID: idExample1,
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
					ResourceID: types.UID{},
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
