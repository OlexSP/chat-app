package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"log/slog"
	"sync"
)

// Config holds the configuration values for the application.
type Config struct {
	HTTPPort       int `env:"HTTP_PORT" env-default:"8080"`
	PostgresConfig PostgresConfig
	MongoConfig    MongoConfig
	KafkaConfig    KafkaConfig
}

// PostgresConfig holds the configuration values for PostgreSQL.
type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     int    `env:"POSTGRES_PORT" env-default:"5432"`
	User     string `env:"POSTGRES_USER" env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"password"`
	DBName   string `env:"POSTGRES_DB" env-default:"chat"`
}

// MongoConfig holds the configuration values for MongoDB.
type MongoConfig struct {
	URI      string `env:"MONGO_URI" env-default:"mongodb://localhost:27017"`
	Database string `env:"MONGO_DATABASE" env-default:"chat"`
}

// KafkaConfig holds the configuration values for Kafka.
type KafkaConfig struct {
	Brokers []string `env:"KAFKA_BROKERS" env-default:"localhost:9092"`
}

var instance *Config
var once sync.Once

// LoadConfig loads the configuration values from environment variables.
func LoadConfig() *Config {
	once.Do(func() {
		log.Println("gather config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "App Notes System"
			description, _ := cleanenv.GetDescription(instance, &helpText)
			slog.Error(description)
			log.Fatal(err)
		}
	})
	return instance
}
