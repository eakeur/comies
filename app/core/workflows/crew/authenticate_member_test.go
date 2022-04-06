package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/app/core/entities/iam/store"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_Authenticate(t *testing.T) {
	t.Parallel()

	fakeID := types.NewUID()

	type (
		fields struct {
			sessions *session.ManagerMock
			stores   *store.ActionsMock
			crew     *crew.ActionsMock
		}

		args struct {
			auth AuthRequest
		}

		test struct {
			name    string
			args    args
			fields  fields
			want    session.Session
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return authenticated session",
			args: args{
				auth: AuthRequest{
					Nickname:       "tovelo@comies",
					Password:       "1922Eakeur!(@@",
					PersistSession: false,
				},
			},
			want: session.Session{
				OperatorID:   fakeID,
				StoreID:      fakeID,
				OperatorName: "Tove Lo",
				Preferences:  types.Preferences{},
			},
			fields: fields{
				sessions: &session.ManagerMock{
					CreateFunc: func(ctx context.Context, op session.Session) (context.Context, session.Session, error) {
						return ctx, op, nil
					},
				},
				stores: &store.ActionsMock{
					ListPreferencesFunc: func(ctx context.Context, storeKey store.Key, modules ...string) (types.Preferences, error) {
						return types.Preferences{}, nil
					},
				},
				crew: &crew.ActionsMock{
					GetMemberWithNicknamesFunc: func(ctx context.Context, operatorNickname string, storeNickname string) (crew.Member, error) {
						return crew.Member{
							Entity: types.Entity{
								ExternalID: fakeID,
							},
							Name:     "Tove Lo",
							FullName: "Tove Lo da Silva",
							Nickname: operatorNickname,
							Password: types.MustCreatePassword("1922Eakeur!(@@"),
							StoreID:  0,
							Store: types.Store{
								StoreID: fakeID,
							},
						}, nil
					},
				},
			},
		},
		{
			name: "should return error for not found member",
			args: args{
				auth: AuthRequest{
					Nickname:       "tovelo@comies",
					Password:       "1922Eakeur!(@@",
					PersistSession: false,
				},
			},
			wantErr: fault.ErrNotFound,
			fields: fields{
				sessions: &session.ManagerMock{
					CreateFunc: func(ctx context.Context, op session.Session) (context.Context, session.Session, error) {
						return ctx, op, nil
					},
				},
				stores: &store.ActionsMock{
					ListPreferencesFunc: func(ctx context.Context, storeKey store.Key, modules ...string) (types.Preferences, error) {
						return types.Preferences{}, nil
					},
				},
				crew: &crew.ActionsMock{
					GetMemberWithNicknamesFunc: func(ctx context.Context, operatorNickname string, storeNickname string) (crew.Member, error) {
						return crew.Member{}, fault.ErrNotFound
					},
				},
			},
		},
		{
			name: "should return error for wrong password",
			args: args{
				auth: AuthRequest{
					Nickname:       "tovelo@comies",
					Password:       "OrWeWillRunAway",
					PersistSession: false,
				},
			},
			wantErr: types.ErrWrongPassword,
			fields: fields{
				sessions: &session.ManagerMock{
					CreateFunc: func(ctx context.Context, op session.Session) (context.Context, session.Session, error) {
						return ctx, op, nil
					},
				},
				stores: &store.ActionsMock{
					ListPreferencesFunc: func(ctx context.Context, storeKey store.Key, modules ...string) (types.Preferences, error) {
						return types.Preferences{}, nil
					},
				},
				crew: &crew.ActionsMock{
					GetMemberWithNicknamesFunc: func(ctx context.Context, operatorNickname string, storeNickname string) (crew.Member, error) {
						return crew.Member{
							Password: types.MustCreatePassword("1922Eakeur!(@@"),
						}, nil
					},
				},
			},
		},
		{
			name: "should return error for invalid nickname",
			args: args{
				auth: AuthRequest{
					Nickname:       "tovelo",
					Password:       "1922Eakeur!(@@",
					PersistSession: false,
				},
			},
			wantErr: crew.ErrInvalidAuthArguments,
			fields: fields{
				sessions: &session.ManagerMock{
					CreateFunc: func(ctx context.Context, op session.Session) (context.Context, session.Session, error) {
						return ctx, op, nil
					},
				},
				stores: &store.ActionsMock{
					ListPreferencesFunc: func(ctx context.Context, storeKey store.Key, modules ...string) (types.Preferences, error) {
						return types.Preferences{}, nil
					},
				},
				crew: &crew.ActionsMock{
					GetMemberWithNicknamesFunc: func(ctx context.Context, operatorNickname string, storeNickname string) (crew.Member, error) {
						return crew.Member{
							Password: types.MustCreatePassword("1922Eakeur!(@@"),
						}, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ses, err := NewWorkflow(c.fields.stores, c.fields.crew, c.fields.sessions).AuthenticateMember(context.Background(), c.args.auth)
			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ses)

		})
	}
}
