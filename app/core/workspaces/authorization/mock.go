package authorization

import "context"

var _ Manager = ManagerMock{}

type ManagerMock struct {
	CreateAccessFunc             func(context.Context) (Access, error)
	CreateAccessResult           Access
	ParseDigestFunc              func(context.Context, string) (Access, error)
	ParseDigestResult            Access
	AssignAccessToContextFunc    func(context.Context, Access) (context.Context, error)
	AssignAccessToContextResult  context.Context
	FetchAccessFromContextFunc   func(context.Context) (Access, error)
	FetchAccessFromContextResult Access
	UpdateAccessWithDigestFunc   func(context.Context, string) (Access, error)
	UpdateAccessWithDigestResult Access
	Error                        error
}

func (m ManagerMock) CreateAccess(ctx context.Context) (Access, error) {
	if m.CreateAccessFunc != nil {
		return m.CreateAccessFunc(ctx)
	}
	return m.CreateAccessResult, m.Error
}

func (m ManagerMock) ParseDigest(ctx context.Context, token string) (Access, error) {
	if m.ParseDigestFunc != nil {
		return m.ParseDigestFunc(ctx, token)
	}
	return m.ParseDigestResult, m.Error
}

func (m ManagerMock) AssignAccessToContext(ctx context.Context, access Access) (context.Context, error) {
	if m.AssignAccessToContextFunc != nil {
		return m.AssignAccessToContextFunc(ctx, access)
	}
	return m.AssignAccessToContextResult, m.Error
}

func (m ManagerMock) FetchAccessFromContext(ctx context.Context) (Access, error) {
	if m.FetchAccessFromContextFunc != nil {
		return m.FetchAccessFromContextFunc(ctx)
	}
	return m.FetchAccessFromContextResult, m.Error
}

func (m ManagerMock) UpdateAccessWithDigest(ctx context.Context, token string) (Access, error) {
	if m.UpdateAccessWithDigestFunc != nil {
		return m.UpdateAccessWithDigestFunc(ctx, token)
	}
	return m.UpdateAccessWithDigestResult, m.Error
}
