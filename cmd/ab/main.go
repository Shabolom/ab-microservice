package main

import (
	"context"
	"net"
	"net/http"
	"os/signal"
	"syscall"

	"ab/internal/di"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	if err := godotenv.Load("./build/local/.env"); err != nil {
		panic(err)
	}

	container := di.New(ctx)
	config := container.Config()

	go func() {
		container.GetWorker().Start(ctx)
	}()

	grpcServer := container.NewAuthGRPCServer()

	metricsMux := container.GetMetrics().MetricsMux()

	go func() {
		addr := ":" + config.Prometheus.Port

		container.Logger().Info(
			"prometheus server started",
			zap.String("addr", addr),
		)

		if err := http.ListenAndServe(addr, metricsMux); err != nil {
			container.Logger().Fatal(
				"prometheus server failed",
				zap.String("addr", addr),
				zap.Error(err),
			)
		}
	}()

	container.Healthcheck()

	go func() {
		lis, err := net.Listen("tcp", config.GRPCPort)
		if err != nil {
			container.Logger().Fatal("failed to listen grpc", zap.Error(err))
		}

		container.Logger().Info(
			"grpc server starting",
			zap.String("addr", container.Config().GRPCPort),
		)

		if err = grpcServer.Serve(lis); err != nil {
			container.Logger().Fatal("grpc server failed", zap.Error(err))
		}
	}()

	<-ctx.Done()
	container.Logger().Info("Shutting down server...")

	container.ShotDown()

	container.Logger().Info("server was shutdown")
}
