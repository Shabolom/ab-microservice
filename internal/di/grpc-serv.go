package di

import (
	authv1 "ab/gen"
	"ab/internal/metrics"
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	grpcHealth "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func (d *DI) NewAuthGRPCServer() *grpc.Server {
	if d.grpcServer != nil {
		return d.grpcServer
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			d.loggingInterceptor(d.Logger()),
			d.metricsInterceptor(d.GetMetrics()),
		),
	)

	authv1.RegisterABExperimentServer(grpcServer, d.GetGRPCHandlers())

	healthServer := health.NewServer()
	healthServer.SetServingStatus(
		"",
		grpcHealth.HealthCheckResponse_SERVING,
	)
	grpcHealth.RegisterHealthServer(grpcServer, healthServer)

	reflection.Register(grpcServer)

	d.grpcServer = grpcServer

	d.Logger().Info("grpc server initialized")

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

		resp, err = handler(ctx, req)

		return resp, err
	}
}

func (d *DI) metricsInterceptor(
	m *metrics.Metrics,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		start := time.Now()

		defer func() {
			m.ObserveGRPCRequest(start, info.FullMethod, err)
		}()

		resp, err = handler(ctx, req)

		return resp, err
	}
}
