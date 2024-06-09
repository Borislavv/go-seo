package server

import (
	"github.com/Borislavv/go-seo/internal/shared/infrastructure/helper"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

const configFile = "http_server.yaml"

type Config struct {
	Port                   string `yaml:"port"`                  // example: ":8080"
	ShutDownTimeoutSeconds int    `yaml:"shutdown_timeout_secs"` // example: 15
}

func (c *Config) Load() *Config {
	path, err := helper.ConfigPath(configFile)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	buff, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(buff, c); err != nil {
		panic(err)
	}

	return c
}
