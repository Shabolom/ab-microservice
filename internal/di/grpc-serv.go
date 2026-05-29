package di

import (
	authv1 "ab/gen"
	"ab/internal/handler/rpctransport"
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (d *DI) NewAuthGRPCServer(logger *zap.Logger, handlers *rpctransport.Handlers) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(d.loggingInterceptor(logger)),
	)

	authv1.RegisterABExperimentServer(grpcServer, handlers)

	reflection.Register(grpcServer)

	logger.Info("grpc server initialized")

	return grpcServer
}

func (d *DI) loggingInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		start := time.Now()

		logger.Info(
			"grpc request started",
			zap.String("method", info.FullMethod),
		)

		defer func() {
			logger.Info(
				"grpc request finished",
				zap.String("method", info.FullMethod),
				zap.Duration("duration", time.Since(start)),
				zap.Error(err),
			)
		}()

		return handler(ctx, req)
	}
}
