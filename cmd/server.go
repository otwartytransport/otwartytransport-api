package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/otwartytransport/otwartytransport-api/internal/server"
)

type Config struct {
	AppPort int `env:"APP_PORT,notEmpty,unset"`
}

func loadConfig(environment string) (Config, error) {
	var cfg Config
	dotenv := fmt.Sprintf(".env.%s", environment)

	if err := godotenv.Load(dotenv); err != nil {
		return cfg, err
	}

	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func main() {
	cfg, err := loadConfig(os.Getenv("APP_ENV"))
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer()
	log.Fatal(srv.Listen(fmt.Sprintf(":%d", cfg.AppPort)))
}
