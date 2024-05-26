package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

// LogrusLogger is decorator under the logrus.Logger which implements additional contextual methods.
// Contextual methods are add extra fields from context.Context.
type LogrusLogger struct {
	*logrus.Logger
	extraFields []string
}

// NewLogger accepts a slice of strings which may contains keys
// which must exists into context.Context when appropriate method calls.
func NewLogger(extraFields []string) (logger *LogrusLogger, closeFunc func()) {
	cfg := new(Config).Load()
	logger = &LogrusLogger{Logger: logrus.New(), extraFields: extraFields}

	logger.SetLevel(cfg.GetLevel())
	logger.SetFormatter(cfg.GetFormat())
	logger.SetReportCaller(true)

	output := cfg.GetOutput()
	logger.SetOutput(output)

	return logger, func() { _ = output.Close() }
}

func (l *LogrusLogger) fieldsFromContext(ctx context.Context) logrus.Fields {
	fields := logrus.Fields{}

	for _, field := range l.extraFields {
		value := ctx.Value(field)
		if value != nil {
			fields[field] = value
		}
	}

	return fields
}

func (l *LogrusLogger) WithField(ctx context.Context, key string, value interface{}) *logrus.Entry {
	return l.Logger.WithFields(l.fieldsFromContext(ctx)).WithField(key, value)
}

func (l *LogrusLogger) WithFields(ctx context.Context, fields logrus.Fields) *logrus.Entry {
	return l.Logger.WithFields(l.fieldsFromContext(ctx)).WithFields(fields)
}

func (l *LogrusLogger) WithError(ctx context.Context, err error) *logrus.Entry {
	return l.Logger.WithFields(l.fieldsFromContext(ctx)).WithError(err)
}

func (l *LogrusLogger) Debug(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Debug(args...)
}

func (l *LogrusLogger) Info(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Info(args...)
}

func (l *LogrusLogger) Print(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Print(args...)
}

func (l *LogrusLogger) Warn(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Warn(args...)
}

func (l *LogrusLogger) Warning(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Warning(args...)
}

func (l *LogrusLogger) Error(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Error(args...)
}

func (l *LogrusLogger) Fatal(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Fatal(args...)
}

func (l *LogrusLogger) Panic(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Panic(args...)
}

func (l *LogrusLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Debugf(format, args...)
}

func (l *LogrusLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Infof(format, args...)
}

func (l *LogrusLogger) Printf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Printf(format, args...)
}

func (l *LogrusLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Warnf(format, args...)
}

func (l *LogrusLogger) Warningf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Warningf(format, args...)
}

func (l *LogrusLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Errorf(format, args...)
}

func (l *LogrusLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Fatalf(format, args...)
}

func (l *LogrusLogger) Panicf(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Panicf(format, args...)
}

func (l *LogrusLogger) Debugln(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Debugln(args...)
}

func (l *LogrusLogger) Infoln(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Infoln(args...)
}

func (l *LogrusLogger) Println(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Println(args...)
}

func (l *LogrusLogger) Warnln(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Warnln(args...)
}

func (l *LogrusLogger) Warningln(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Warningln(args...)
}

func (l *LogrusLogger) Errorln(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Errorln(args...)
}

func (l *LogrusLogger) Fatalln(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Fatalln(args...)
}

func (l *LogrusLogger) Panicln(ctx context.Context, args ...interface{}) {
	l.Logger.WithFields(l.fieldsFromContext(ctx)).Panicln(args...)
}
