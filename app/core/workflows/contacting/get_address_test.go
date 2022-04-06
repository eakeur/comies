package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_GetAddress(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

	type (
		args struct {
			targetID  types.UID
			addressID types.UID
		}

		opts struct {
			contacts *contacting.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    contacting.Address
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return address",
			args: args{
				targetID:  fakeID,
				addressID: fakeID,
			},
			want: contacting.Address{},
			opts: opts{
				contacts: &contacting.ActionsMock{
					GetAddressFunc: func(ctx context.Context, target types.UID, addressID types.UID) (contacting.Address, error) {
						return contacting.Address{}, nil
					},
				},
			},
		},
		{
			name: "should return address not found",
			args: args{
				targetID:  fakeID,
				addressID: fakeID,
			},
			wantErr: fault.ErrNotFound,
			opts: opts{
				contacts: &contacting.ActionsMock{
					GetAddressFunc: func(ctx context.Context, target types.UID, addressID types.UID) (contacting.Address, error) {
						return contacting.Address{}, fault.ErrNotFound
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.contacts)
			got, err := wf.GetAddress(ctx, c.args.targetID, c.args.addressID)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

		})
	}
}
