package logger

import (
	"github.com/Borislavv/go-seo/internal/shared/infrastructure/helper"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"path/filepath"
)

const configFile = "logger.yaml"

type Config struct {
	Output string `yaml:"output"`
	Format string `yaml:"format"`
	Level  string `yaml:"level"`
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

func (c *Config) GetOutput() *os.File {
	if c.Output == "stdout" {
		return os.Stdout
	}

	path := ""
	if c.Output == "" {
		path = "/dev/null"
	} else {
		fpath, err := helper.LogsPath(c.Output)
		if err != nil {
			panic(err)
		}
		path = fpath
	}

	if _, err := os.ReadDir(filepath.Dir(path)); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func (c *Config) GetFormat() logrus.Formatter {
	switch c.Format {
	case "text":
		return &logrus.TextFormatter{}
	default:
		return &logrus.JSONFormatter{}
	}
}

func (c *Config) GetLevel() logrus.Level {
	switch c.Level {
	case "info":
		return logrus.InfoLevel
	case "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}
