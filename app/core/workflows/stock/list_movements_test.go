package stock

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/stocking/stock"
	"gomies/pkg/sdk/types"
	"testing"
	"time"
)

func TestWorkflow_ListMovements(t *testing.T) {
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
			want    []stock.Movement
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should list period successfully",
			args: args{
				filter: stock.Filter{
					ResourceID: idExample1,
					FinalDate:  time.Now(),
				},
			},
			want: []stock.Movement{},
			opts: opts{
				stocks: &stock.ActionsMock{
					ListMovementsFunc: func(ctx context.Context, filter stock.Filter) ([]stock.Movement, error) {
						return []stock.Movement{}, nil
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
			want:    []stock.Movement{},
			opts: opts{
				stocks: &stock.ActionsMock{
					ListMovementsFunc: func(ctx context.Context, filter stock.Filter) ([]stock.Movement, error) {
						return []stock.Movement{}, nil
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
			lis, err := wf.ListMovements(ctx, c.args.filter)

			assert.Equal(t, c.want, lis)
			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
