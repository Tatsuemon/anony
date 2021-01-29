package middleware

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor(authFunc AuthFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, err := authFunc(ctx, info.FullMethod)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		res, err := handler(newCtx, req)
		if err != nil {
			return nil, err
		}
		// log.Info("%+v\n", res)
		return res, nil
	}
}
