package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Values struct {
	Address     string       `envconfig:"SERVER_ADDRESS"`
	Hostname    string       `envconfig:"BASE_URL"`
	DBConn      *DBValues    `envconfig:"DB"`
	RedisConn   *RedisValues `envconfig:"REDIS"`
	StorageType string       `envconfig:"STORAGE_TYPE"`
}

type DBValues struct {
	DBUser    string `envconfig:"USER"`
	DBPass    string `envconfig:"PASSWORD"`
	DBName    string `envconfig:"NAME"`
	DBHost    string `envconfig:"HOST"`
	DBPort    string `envconfig:"PORT"`
	DBSslMode string `envconfig:"SSLMODE"`
}

type RedisValues struct {
	RedisAddr string `envconfig:"ADDR"`
	RedisPass string `envconfig:"PASSWORD"`
}

func LoadConfig(getFlags bool) (*Values, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file:", err)
	}

	cfg := &Values{}
	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	address := new(string)
	hostname := new(string)
	dbconn := new(*DBValues)
	redisconn := new(*RedisValues)
	stype := new(string)

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

	if cfg.DBConn == nil {
		if *dbconn == nil {
			*dbconn = &DBValues{
				DBUser:    "postgres",
				DBPass:    "secret",
				DBHost:    "localhost",
				DBName:    "db-store",
				DBPort:    "5433",
				DBSslMode: "disable",
			}
		}
		cfg.DBConn = *dbconn
	}

	if cfg.RedisConn == nil {
		if *redisconn == nil {
			*redisconn = &RedisValues{
				RedisAddr: "db:6379",
				RedisPass: "",
			}
		}
		cfg.RedisConn = *redisconn
	}

	if cfg.StorageType == "" {
		if *stype == "" {
			*stype = "pg"
		}
		cfg.StorageType = *stype
	}

	return cfg, nil
}

func BuildDSN(cfg *Values) string {

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBConn.DBUser,
		cfg.DBConn.DBPass,
		cfg.DBConn.DBHost,
		cfg.DBConn.DBPort,
		cfg.DBConn.DBName,
		cfg.DBConn.DBSslMode,
	)
}
