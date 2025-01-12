package proto

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log/slog"
)

// UnaryClientErrorInterceptor - interceptor для unary RPC
func UnaryClientErrorInterceptor(
	log *slog.Logger,
) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			log.Error("gRPC Client Error in method %s: %v", method, err)
			return status.Errorf(status.Code(err), "Custom error: %v", err)
		}
		return nil
	}
}
