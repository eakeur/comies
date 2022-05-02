package session

import (
	"context"
	"gomies/app/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSession_Delegate(t *testing.T) {
	t.Parallel()

	fakeID := types.ID(1)

	type (
		args struct {
			operation string
			store     *types.Store
			history   *types.History
		}

		test struct {
			name        string
			args        args
			session     Session
			wantStore   types.Store
			wantHistory types.History
		}
	)

	cases := []test{
		{
			name: "should fill all fields",
			args: args{
				operation: "Workflows.Product.CreateProduct",
				store:     &types.Store{},
				history:   &types.History{},
			},
			session: Session{
				OperatorID: fakeID,
				StoreID:    fakeID,
			},
			wantStore:   types.Store{StoreID: fakeID},
			wantHistory: types.History{Operation: "Workflows.Product.CreateProduct", By: fakeID},
		},
		{
			name: "should fill store fields only",
			args: args{
				operation: "Workflows.Product.CreateProduct",
				store:     &types.Store{},
				history:   nil,
			},
			session: Session{
				OperatorID: fakeID,
				StoreID:    fakeID,
			},
			wantStore: types.Store{StoreID: fakeID},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			c.session.Delegate(c.args.operation, c.args.store, c.args.history)

			assert.Equal(t, c.wantStore.StoreID, c.args.store.StoreID)

			if c.args.store != nil {
				assert.Equal(t, c.wantStore.StoreID, c.args.store.StoreID)
			}

			if c.args.history != nil {
				assert.Equal(t, c.wantHistory.By, c.args.history.By)
				assert.Equal(t, c.wantHistory.Operation, c.args.history.Operation)
			}
		})
	}
}

func TestSession_WithContext(t *testing.T) {
	t.Parallel()

	fakeID := types.ID(1)

	type test struct {
		name    string
		ctx     context.Context
		session *Session
		wantErr error
	}

	cases := []test{
		{
			name: "should send and retrieve session in context",
			ctx:  context.Background(),
			session: &Session{
				OperatorID: fakeID,
				StoreID:    fakeID,
			},
		},
		{
			name:    "should fail with ErrNoSession",
			ctx:     context.Background(),
			session: nil,
			wantErr: ErrNoSession,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			if c.session != nil {
				ctx = c.session.WithContext(c.ctx)
			}

			ses, err := FromContext(ctx)

			assert.ErrorIs(t, err, c.wantErr)

			if err == nil && c.session != nil {
				assert.Equal(t, *c.session, ses)
			}

		})
	}
}
