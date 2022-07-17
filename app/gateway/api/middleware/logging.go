package middleware

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func (m middleware) Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		at := time.Now()

		res, err := handler(ctx, req)
		log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
			info.FullMethod,
			time.Since(at),
			err)

		return res, err
	}
}
