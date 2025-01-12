package proto

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryServerErrorInterceptor - interceptor для unary RPC, перехватывает и обрабатывает ошибки
func UnaryServerErrorInterceptor(
	log *slog.Logger,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			log.Error("gRPC Error in method %s: %v", info.FullMethod, err)
			// Здесь можно кастомизировать статус ошибки или отправить её "как есть"
			// Например, можно оставить только её публичный статус
			return nil, status.Convert(err).Err()
		}
		return resp, nil
	}
}
