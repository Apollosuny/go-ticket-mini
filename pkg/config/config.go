package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppEnv string
	TicketServicePort string
	PostgresDSN string
}

func Load() (*Config, error) {
	cfg := &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		TicketServicePort: getEnv("TICKET_SERVICE_PORT", "9091"),
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