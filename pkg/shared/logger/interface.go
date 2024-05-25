package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	logrus.StdLogger
	logrus.FieldLogger
	logrus.Ext1FieldLogger
	GetRequestCtx() context.Context
	SetRequestCtx(context.Context)
}
