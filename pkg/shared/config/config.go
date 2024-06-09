package config

import (
	"log"
	"os"
)

const (
	ProdEnv = "prod"
	DevEnv  = "dev"
	TestEnv = "test"
)

var cfg *Config

type Config struct {
	appEnv string `env:"APP_ENV"`
}

func Load() {
	cfg = new(Config)
	cfg.appEnv = os.Getenv("APP_ENV")
	if cfg.appEnv == "" {
		log.Panic("APP_ENV environment variable must be defined")
	} else {
		log.Printf("APP_ENV=%v\n", cfg.appEnv)
	}
}

func Get() *Config {
	return cfg
}

func (c *Config) IsProd() bool {
	return c.appEnv == ProdEnv
}

func (c *Config) IsDev() bool {
	return c.appEnv == DevEnv
}

func (c *Config) IsTest() bool {
	return c.appEnv == TestEnv
}
