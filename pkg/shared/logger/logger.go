package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() (logger *logrus.Logger, closeFunc func()) {
	cfg := new(Config).Load()
	lgr := logrus.New()
	lgr.SetLevel(cfg.GetLevel())

	lgr.SetFormatter(cfg.GetFormat())
	lgr.SetReportCaller(true)

	output := cfg.GetOutput()
	lgr.SetOutput(output)

	return lgr, func() { _ = output.Close() }
}
