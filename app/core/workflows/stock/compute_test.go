package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/pkg/sdk/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_Compute(t *testing.T) {
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
			want    types.Quantity
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should compute period successfully",
			args: args{
				filter: stock.Filter{
					ResourceID: fakeID,
					FinalDate:  time.Now(),
				},
			},
			want: 1000,
			opts: opts{
				stocks: &stock.ActionsMock{
					ComputeFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						return 1000, nil
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
					ComputeFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						return 1000, nil
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
			calc, err := wf.Compute(ctx, c.args.filter)

			assert.Equal(t, c.want, calc)
			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}

func TestWorkflow_ComputeSome(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

	type (
		args struct {
			resourcesIDs []types.UID
			filter       stock.Filter
		}

		opts struct {
			stocks *stock.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    []types.Quantity
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should compute period successfully",
			args: args{
				resourcesIDs: []types.UID{
					fakeID, fakeID,
				},
				filter: stock.Filter{
					FinalDate: time.Now(),
				},
			},
			want: []types.Quantity{1000, 1000},
			opts: opts{
				stocks: &stock.ActionsMock{
					ComputeSomeFunc: func(ctx context.Context, filter stock.Filter, resourceID ...types.UID) ([]types.Quantity, error) {
						return []types.Quantity{1000, 1000}, nil
					},
				},
			},
		},
		{
			name: "should fail because resourceID is invalid",
			args: args{
				resourcesIDs: []types.UID{
					fakeID, "",
				},
				filter: stock.Filter{
					FinalDate: time.Now(),
				},
			},
			wantErr: stock.ErrMissingResourceID,
			want:    []types.Quantity{},
			opts: opts{
				stocks: &stock.ActionsMock{
					ComputeSomeFunc: func(ctx context.Context, filter stock.Filter, resourceID ...types.UID) ([]types.Quantity, error) {
						return []types.Quantity{1000, 1000}, nil
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
			calc, err := wf.ComputeSome(ctx, c.args.filter, c.args.resourcesIDs...)

			assert.Equal(t, c.want, calc)
			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
