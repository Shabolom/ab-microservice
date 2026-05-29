package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type PgDB struct {
	Host          string `envconfig:"POSTGRES_HOST"`
	Port          string `envconfig:"POSTGRES_PORT"`
	User          string `envconfig:"POSTGRES_USER"`
	Password      string `envconfig:"POSTGRES_PASSWORD"`
	Database      string `envconfig:"POSTGRES_DB"`
	SSLMode       string `envconfig:"POSTGRES_SSLMODE"`
	MaxConnection int    `envconfig:"POSTGRES_MAX_CONNECTION" default:"10"`
	MinConnection int    `envconfig:"POSTGRES_MIN_CONNECTION" default:"0"`
}
type Config struct {
	ServiceName    string `envconfig:"APP_NAME"`
	Debug          bool   `envconfig:"APP_DEBUG"`
	GRPCPort       string `envconfig:"APP_GRPC_ADDRESS"`
	Secret         string `envconfig:"APP_SECRET"`
	ResendAppKey   string `envconfig:"RESEND_API_KEY"`
	Kafka          Kafka
	KafkaSerialize KafkaSerialize
	PostgresDB     PgDB
}

type KafkaSerialize struct {
	Host string `envconfig:"SCHEMA_REGISTRY_HOST"`
	Port string `envconfig:"SCHEMA_REGISTRY_PORT"`
}
type Kafka struct {
	Brokers []string `envconfig:"KAFKA_BROKERS"`
	Topic   string   `envconfig:"KAFKA_TOPIC"`
	GroupID string   `envconfig:"KAFKA_GROUP_ID"`
}

func FromEnv() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("error while parse env config | %w", err)
	}

	return cfg, nil
}

func (c *Config) SchemaRegisterDSN() string {
	return fmt.Sprintf("http://%v:%v",
		c.KafkaSerialize.Host,
		c.KafkaSerialize.Port)
}

func (c *Config) PostgresDBURL() string {
	pgURL := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v",
		c.PostgresDB.User,
		c.PostgresDB.Password,
		c.PostgresDB.Host,
		c.PostgresDB.Port,
		c.PostgresDB.Database,
	)

	return pgURL
}
