package proto

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
)

// UnaryClientErrorInterceptor - interceptor для unary RPC
func UnaryClientErrorInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Вызов RPC метода
		err := invoker(ctx, method, req, reply, cc, opts...)

		// Если вызов вернул ошибку, обработайте её
		if err != nil {
			// Логирование ошибки
			log.Printf("gRPC Client Error in method %s: %v", method, err)

			// Пример: можно обернуть ошибку в пользовательский формат
			return status.Errorf(status.Code(err), "Custom error: %v", err)
		}

		// Если ошибки нет, вернуть nil
		return nil
	}
}

// StreamClientErrorInterceptor - interceptor для stream RPC
func StreamClientErrorInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		// Вызов stream метода
		clientStream, err := streamer(ctx, desc, cc, method, opts...)

		// Если stream вызов вернул ошибку, обработайте её
		if err != nil {
			// Логирование ошибки
			log.Printf("gRPC Client Stream Error in method %s: %v", method, err)

			// Пример: можно обернуть ошибку в пользовательский формат
			return nil, status.Errorf(status.Code(err), "Custom error: %v", err)
		}

		// Если ошибки нет, вернуть stream
		return clientStream, nil
	}
}
