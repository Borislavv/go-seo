package helper

import (
	"os"
	"strings"
)

const (
	LogsDir = "/var/log"
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

func LogsPath(additionalPath string) (path string, err error) {
	return Path(LogsDir + string(os.PathSeparator) + additionalPath)
}
