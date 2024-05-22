package helper

import (
	"github.com/Borislavv/go-seo/pkg/shared/config"
	"os"
	"strings"
)

const (
	LogsDir       = "/var/log"
	ConfigDevDir  = "/cfg/dev"
	ConfigProdDir = "/cfg/prod"
	ConfigTestDir = "/cfg/dev"
)

var (
	appEnv = os.Getenv("APP_ENV")
)

func Path(additionalPath string) (path string, err error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}

	pathSeparator := string(os.PathSeparator)
	path = root + pathSeparator + additionalPath
	path = strings.ReplaceAll(path, pathSeparator+pathSeparator, pathSeparator)

	return path, nil
}

func ConfigPath(additionalPath string) (path string, err error) {
	path = ConfigDevDir
	if config.Get().IsProd() {
		path = ConfigProdDir
	} else if config.Get().IsTest() {
		path = ConfigTestDir
	}

	return Path(path + string(os.PathSeparator) + additionalPath)
}

func LogsPath(additionalPath string) (path string, err error) {
	return Path(LogsDir + string(os.PathSeparator) + additionalPath)
}
