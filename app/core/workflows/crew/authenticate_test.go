package crew

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gomies/app/core/entities/crew"
	"gomies/app/core/managers/session"
	"testing"
)

func TestWorkflow_Authenticate(t *testing.T) {

	t.Parallel()

	type test struct {
		name    string
		args    crew.AuthRequest
		opts    workflow
		want    session.Session
		wantErr error
	}

	cases := []test{
		{
			name: "should return successful auth",
			args: crew.AuthRequest{
				Nickname:       "igor@avis",
				Password:       "12345678",
				PersistSession: false,
			},
			want: session.Session{},
			opts: workflow{
				crew: &crew.ActionsMock{
					GetWithNickFunc: func(ctx context.Context, nick string, store string) (crew.Operator, error) {
						arr, _ := bcrypt.GenerateFromPassword([]byte("12345678"), 1)
						return crew.Operator{
							FirstName:     "Igor",
							Nick:     nick,
							Password: string(arr),
						}, nil
					},
					UpdateFunc: func(ctx context.Context, operator crew.Operator) error {
						return nil
					},
				},
				sessions: &session.ManagerMock{
					CreateFunc: func(ctx context.Context, op session.Session) (context.Context, session.Session, error) {
						return ctx, session.Session{}, nil
					},
				},
			},
		},
		{
			name: "should fail auth because password is wrong",
			args: crew.AuthRequest{
				Nickname:       "igor@avis",
				Password:       "123456789",
				PersistSession: false,
			},
			want:    session.Session{},
			wantErr: crew.ErrWrongPassword,
			opts: workflow{
				crew: &crew.ActionsMock{
					GetWithNickFunc: func(ctx context.Context, nick string, store string) (crew.Operator, error) {
						arr, _ := bcrypt.GenerateFromPassword([]byte("12345678"), 1)
						return crew.Operator{
							FirstName:     "Igor",
							Nick:     nick,
							Password: string(arr),
						}, nil
					},
				},
			},
		},
		{
			name: "should fail auth because operator does not exist",
			args: crew.AuthRequest{
				Nickname:       "igor@avis",
				Password:       "123456789",
				PersistSession: false,
			},
			want:    session.Session{},
			wantErr: crew.ErrNotFound,
			opts: workflow{
				crew: &crew.ActionsMock{
					GetWithNickFunc: func(ctx context.Context, nick string, store string) (crew.Operator, error) {
						return crew.Operator{}, crew.ErrNotFound
					},
				},
			},
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			wf := NewWorkflow(tc.opts.stores, tc.opts.crew, tc.opts.transactions, tc.opts.sessions)

			gotSession, gotErr := wf.Authenticate(context.Background(), tc.args)

			assert.ErrorIs(t, gotErr, tc.wantErr)
			assert.Equal(t, tc.want, gotSession)
		})
	}
}
