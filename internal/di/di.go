package di

import (
	inMemmoryCashe "ab/internal/in-memory-cache"
	KafkaProducer "ab/internal/kafka-producer"
	"ab/internal/metrics"
	"context"
	"fmt"

	"ab/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type DI struct {
	config *config.Config
	logger *zap.Logger

	kafka *KafkaProducer.Kafka

	inMemoryCache *inMemmoryCashe.RawExperimentSessionStorage

	pgConn *pgxpool.Pool

	grpcServer *grpc.Server

	metrics *metrics.Metrics
}

func New(ctx context.Context) *DI {
	_ = ctx
	return &DI{}
}

func (d *DI) Config() *config.Config {
	if d.config != nil {
		return d.config
	}

	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Errorf("config from env: %w", err))
	}

	d.config = cfg
	return d.config
}

func (d *DI) Logger() *zap.Logger {
	if d.logger != nil {
		return d.logger
	}

	var logger *zap.Logger
	var err error

	if d.Config().Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(fmt.Errorf("create logger: %w", err))
	}

	logger = logger.With(
		zap.String("service", d.Config().ServiceName),
		zap.Bool("debug", d.Config().Debug),
	)

	d.logger = logger
	_ = zap.ReplaceGlobals(logger)

	return d.logger
}

func (d *DI) ShotDown() {
	d.Logger().Info("Shutting kafka...")
	err := d.kafka.Close()
	if err != nil {
		d.Logger().Info("Failed to shut down kafka", zap.Error(err))
	}
	d.Logger().Info("Kafka closed")

	d.Logger().Info("pgDB shutting down...")
	d.pgConn.Close()
	d.Logger().Info("pgConn closed")

	d.Logger().Info("Grpc server shutting down...")
	d.grpcServer.Stop()
	d.Logger().Info("Grpc server stoped")
}
