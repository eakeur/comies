package authorization

import "context"

type Manager interface {
	CreateAccess(context.Context) (Access, error)

	ParseDigest(context.Context, string) (Access, error)

	AssignAccessToContext(context.Context, Access) (context.Context, error)

	FetchAccessFromContext(context.Context) (Access, error)

	UpdateAccessWithDigest(context.Context, string) (Access, error)
}
