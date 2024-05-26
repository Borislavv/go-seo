package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	WithField(ctx context.Context, key string, value interface{}) *logrus.Entry
	WithFields(ctx context.Context, fields logrus.Fields) *logrus.Entry
	WithError(ctx context.Context, err error) *logrus.Entry

	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Printf(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Warningf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})

	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Print(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})
	Panic(ctx context.Context, args ...interface{})

	Debugln(ctx context.Context, args ...interface{})
	Infoln(ctx context.Context, args ...interface{})
	Println(ctx context.Context, args ...interface{})
	Warnln(ctx context.Context, args ...interface{})
	Warningln(ctx context.Context, args ...interface{})
	Errorln(ctx context.Context, args ...interface{})
	Fatalln(ctx context.Context, args ...interface{})
	Panicln(ctx context.Context, args ...interface{})
}
