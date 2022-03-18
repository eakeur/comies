package stock

//
//import (
//	"context"
//	"github.com/stretchr/testify/assert"
//	"gomies/pkg/catalog/core/entities/product"
//	product2 "gomies/pkg/catalog/core/workflows/product"
//	"gomies/pkg/sdk/fault"
//	"gomies/pkg/sdk/tests"
//	"gomies/pkg/sdk/types"
//	"gomies/pkg/stocking/core/entities/stock"
//	"testing"
//)
//
//func TestWorkflow_RemoveFromStock(t *testing.T) {
//
//	const operation = "Workflows.Product.RemoveIngredient"
//	t.Parallel()
//
//	ctx := tests.WorkflowContext(product2.idExample1, product2.idExample2)
//	managers := tests.Managers()
//
//	type (
//		args struct {
//			key product.Key
//			id  types.UID
//		}
//
//		opts struct {
//			stocks   *stock.ActionsMock
//			products *product.ActionsMock
//		}
//
//		test struct {
//			name    string
//			args    args
//			opts    opts
//			wantKey product.Key
//			wantErr error
//		}
//	)
//
//	cases := []test{
//		{
//			name: "should remove movement",
//			args: args{
//				key: product.Key{ID: product2.idExample1},
//				id:  product2.idExample3,
//			},
//			wantKey: product.Key{ID: product2.idExample1, Store: types.Store{StoreID: product2.idExample2}},
//			opts: opts{
//				products: &product.ActionsMock{
//					GetFunc: func(ctx context.Context, key product.Key) (product.Product, error) {
//						return product.Product{}, nil
//					},
//				},
//				stocks: &stock.ActionsMock{
//					RemoveFromStockFunc: func(ctx context.Context, target types.UID, movementID types.UID) error {
//						return nil
//					},
//				},
//			},
//		},
//		{
//			name: "should return product not found",
//			args: args{
//				key: product.Key{ID: product2.idExample1},
//				id:  product2.idExample3,
//			},
//			wantErr: fault.ErrNotFound,
//			wantKey: product.Key{ID: product2.idExample1, Store: types.Store{StoreID: product2.idExample2}},
//			opts: opts{
//				products: &product.ActionsMock{
//					GetFunc: func(ctx context.Context, key product.Key) (product.Product, error) {
//						return product.Product{}, fault.ErrNotFound
//					},
//				},
//				stocks: &stock.ActionsMock{
//					RemoveFromStockFunc: func(ctx context.Context, target types.UID, movementID types.UID) error {
//						return nil
//					},
//				},
//			},
//		},
//	}
//
//	for _, c := range cases {
//		c := c
//
//		t.Run(c.name, func(t *testing.T) {
//			t.Parallel()
//
//			wf := product2.NewWorkflow(c.opts.products, c.opts.stocks, nil, managers.Transactions)
//			err := wf.RemoveFromStock(ctx, c.args.key, c.args.id)
//
//			assert.ErrorIs(t, err, c.wantErr)
//
//			if err == nil && c.wantErr == nil {
//				assert.Equal(t, c.wantKey, c.opts.products.GetCalls()[0].Key)
//			}
//
//		})
//	}
//}
