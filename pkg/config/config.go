package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string
	APIGatewayPort string
	TicketServicePort string
	TicketServiceGRPCAddr string
	PostgresDSN string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("err: %s", err)
		log.Println("No .env file found or error loading .env file")
	}

	cfg := &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		APIGatewayPort: getEnv("API_GATEWAY_PORT", "8080"),
		TicketServicePort: getEnv("TICKET_SERVICE_PORT", "9091"),
		TicketServiceGRPCAddr: getEnv("TICKET_SERVICE_GRPC_ADDR", "localhost:9091"),
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.PostgresDSN == "" {
		return fmt.Errorf("POSTGRES_DSN is required")
	}
	return nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}