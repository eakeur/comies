package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_GetPhone(t *testing.T) {
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
			want    contacting.Phone
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return phone",
			args: args{
				targetID: fakeID,
				phoneID:  fakeID,
			},
			want: contacting.Phone{},
			opts: opts{
				contacts: &contacting.ActionsMock{
					GetPhoneFunc: func(ctx context.Context, target types.UID, id types.UID) (contacting.Phone, error) {
						return contacting.Phone{}, nil
					},
				},
			},
		},
		{
			name: "should return phone not found",
			args: args{
				targetID: fakeID,
				phoneID:  fakeID,
			},
			wantErr: fault.ErrNotFound,
			opts: opts{
				contacts: &contacting.ActionsMock{
					GetPhoneFunc: func(ctx context.Context, target types.UID, id types.UID) (contacting.Phone, error) {
						return contacting.Phone{}, fault.ErrNotFound
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
			got, err := wf.GetPhone(ctx, c.args.targetID, c.args.phoneID)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

		})
	}
}
