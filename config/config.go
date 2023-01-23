package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Host        string `env:"HOST"         envDefault:"0.0.0.0"`
	Port        int    `env:"PORT"         envDefault:"3000"`
	Debug       bool   `env:"DEBUG"        envDefault:"false"`
	DatabaseUrl string `env:"DATABASE_URL"`
}

var Config *config

func Initalize() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse config: %+v\n", err)
	}
	log.Infof("env var: %+v\n", cfg)
	Config = &cfg
}
