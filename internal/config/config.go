package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	ServiceName string
	Debug       bool
	Port        string
	Secret      string
	Postgres    PostgresConfig
	Kafka       KafkaConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
	SSLMode  string
	DSN      string
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
	GroupID string
}

func FromEnv() (*Config, error) {
	cfg := &Config{
		ServiceName: getEnv("APP_NAME", "app"),
		Debug:       getEnvBool("APP_DEBUG", true),
		Port:        getEnv("APP_PORT", "8080"),
		Secret:      getEnv("APP_SECRET", "secret"),

		Postgres: PostgresConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			DB:       getEnv("POSTGRES_DB", "postgres"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "postgres"),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
			DSN:      os.Getenv("POSTGRES_DSN"),
		},

		Kafka: KafkaConfig{
			Brokers: splitAndClean(getEnv("KAFKA_BROKERS", "localhost:9092")),
			Topic:   getEnv("KAFKA_TOPIC", "example-topic"),
			GroupID: getEnv("KAFKA_GROUP_ID", "example-group"),
		},
	}

	return cfg, nil
}

func (c *Config) DatabaseURL() string {
	if c.Postgres.DSN != "" {
		return c.Postgres.DSN
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.DB,
		c.Postgres.SSLMode,
	)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}

	return parsed
}
