package main

import (
	"context"
	"net"
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

	go func() {
		container.GetWorker().Start(ctx)
	}()

	grpcServer := container.NewAuthGRPCServer()

	go func() {
		lis, err := net.Listen("tcp", container.Config().GRPCPort)
		if err != nil {
			container.Logger().Fatal(error.Error(err))
		}

		container.Logger().Info(
			"grpc server started",
			zap.String("port", container.Config().GRPCPort),
		)
		if err = grpcServer.Serve(lis); err != nil {
			container.Logger().Fatal(error.Error(err))
		}
	}()

	<-ctx.Done()
	container.Logger().Info("Shutting down server...")

	container.ShotDown()

	container.Logger().Info("server was shutdown")
}
