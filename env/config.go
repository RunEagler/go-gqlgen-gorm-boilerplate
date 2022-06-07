package env

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type config struct {
	DbHost string `env:"DB_HOST"`
	DbPort string `env:"DB_PORT"`
	DbName string `env:"DB_NAME"`
	DbUser string `env:"DB_USER"`
}

var Config config

func SetUpEnv() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	Config = cfg
}
