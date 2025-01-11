package proto

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryServerErrorInterceptor - interceptor для unary RPC, перехватывает и обрабатывает ошибки
func UnaryServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// Вызов основного handler (реализации RPC-метода)
		resp, err = handler(ctx, req)

		// Если метод вернул ошибку, обработайте её
		if err != nil {
			// Логирование ошибки
			log.Printf("gRPC Error in method %s: %v", info.FullMethod, err)

			// Здесь можно кастомизировать статус ошибки или отправить её "как есть"
			// Например, можно оставить только её публичный статус
			return nil, status.Convert(err).Err()
		}

		// Возвращаем без изменений, если нет ошибки
		return resp, nil
	}
}

// StreamServerErrorInterceptor - interceptor для stream RPC
func StreamServerErrorInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		// Вызов основного stream-обработчика
		err := handler(srv, ss)

		// Если stream вернул ошибку, обработайте её
		if err != nil {
			// Логирование ошибки
			log.Printf("gRPC Stream Error in method %s: %v", info.FullMethod, err)

			// Возвратим только её публичный статус
			return status.Convert(err).Err()
		}

		// Возвратим nil, если ошибок нет
		return nil
	}
}
