package di

import (
	KafkaProducer "ab/internal/kafka-producer"
	"context"
	"fmt"

	"ab/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type DI struct {
	config *config.Config
	logger *zap.Logger

	kafka *KafkaProducer.Kafka

	pgConn *pgxpool.Pool
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
