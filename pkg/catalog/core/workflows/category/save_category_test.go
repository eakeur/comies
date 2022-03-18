package category

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_SaveCategory(t *testing.T) {
	const operation = "Workflows.Category.SaveCategory"
	t.Parallel()

	ctx := tests.WorkflowContext(idExample1, idExample2)
	managers := tests.Managers()

	type (
		args struct {
			category category.Category
		}

		opts struct {
			categories *category.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    category.Category
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return category created",
			args: args{
				category: category.Category{
					Code: "PRD",
					Name: "Product",
				},
			},
			want: category.Category{
				Entity: types.Entity{
					History: types.History{
						By:        idExample1,
						Operation: operation,
					},
				},
				Code:  "PRD",
				Name:  "Product",
				Store: types.Store{StoreID: idExample2},
			},
			opts: opts{
				categories: &category.ActionsMock{
					SaveFunc: func(ctx context.Context, prd category.Category, flag ...types.WritingFlag) (category.Category, error) {
						return prd, nil
					},
				},
			},
		},
		{
			name: "should return error for invalid code",
			args: args{
				category: category.Category{
					Code: "P",
					Name: "Product",
				},
			},
			wantErr: category.ErrInvalidCode,
			opts: opts{
				categories: &category.ActionsMock{
					SaveFunc: func(ctx context.Context, prd category.Category, flag ...types.WritingFlag) (category.Category, error) {
						return prd, nil
					},
				},
			},
		},
		{
			name: "should return error for invalid name",
			args: args{
				category: category.Category{
					Code: "PRD",
					Name: "-",
				},
			},
			wantErr: category.ErrInvalidName,
			opts: opts{
				categories: &category.ActionsMock{
					SaveFunc: func(ctx context.Context, prd category.Category, flag ...types.WritingFlag) (category.Category, error) {
						return prd, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.categories, managers.Transactions)
			ingredient, err := wf.SaveCategory(ctx, c.args.category)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ingredient)

		})
	}

}
