package middleware

import (
	"context"
	"google.golang.org/grpc"
)

func (m middleware) Transaction() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = m.managers.Transactions.Begin(ctx)
		defer m.managers.Transactions.End(ctx)

		return handler(ctx, req)
	}
}
