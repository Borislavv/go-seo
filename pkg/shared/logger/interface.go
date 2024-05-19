package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	logrus.StdLogger
	logrus.FieldLogger
	logrus.Ext1FieldLogger
}
