package redis

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	client *redis.Client
	logger *zap.Logger
}

func New(client *redis.Client, logger *zap.Logger) *Redis {
	return &Redis{
		client: client,
		logger: logger,
	}
}
