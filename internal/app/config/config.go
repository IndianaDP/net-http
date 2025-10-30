// internal/config/config.go
package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
)

type Values struct {
	Address  string `env:"SERVER_ADDRESS"`
	Hostname string `env:"BASE_URL"`
	UUID     string `env:"UUID"`
}

func LoadConfig(getFlags bool) (*Values, error) {

	var cfg Values
	address := new(string)
	hostname := new(string)

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}

	if getFlags {
		address = flag.String("a", "", "Server address")
		hostname = flag.String("h", "", "Base hostname")
		flag.Parse()
	}

	if cfg.Address == "" {
		if *address == "" {
			*address = fmt.Sprintf(`:%d`, 8080)
		}

		cfg.Address = *address
	}
	if cfg.Hostname == "" {
		if *hostname == "" {
			*hostname = "http://localhost:8080"
		}

		cfg.Hostname = *hostname
	}

	cfg.UUID = "i2o3hgo3ihg"
	return &cfg, nil
}
