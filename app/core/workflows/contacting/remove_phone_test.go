package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_RemovePhoneTest(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

	type (
		args struct {
			targetID types.UID
			phoneID  types.UID
		}

		opts struct {
			contacts *contacting.ActionsMock
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
			name: "should return nil error",
			args: args{
				targetID: fakeID,
				phoneID:  fakeID,
			},
			opts: opts{
				contacts: &contacting.ActionsMock{
					RemovePhonesFunc: func(ctx context.Context, target types.UID, ids ...types.UID) error {
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

			wf := NewWorkflow(c.opts.contacts)
			err := wf.RemovePhone(ctx, c.args.targetID, c.args.phoneID)

			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
