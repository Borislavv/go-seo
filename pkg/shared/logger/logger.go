package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	requestCtx context.Context
	processor  Processor
	*logrus.Logger
}

func NewLogger(defaultCtx context.Context, processors []FieldsProcessor) (logger *LogrusLogger, closeFunc func()) {
	cfg := new(Config).Load()
	lgr := &LogrusLogger{requestCtx: defaultCtx, Logger: logrus.New()}
	lgr.SetLevel(cfg.GetLevel())

	lgr.SetFormatter(cfg.GetFormat())
	lgr.SetReportCaller(true)

	output := cfg.GetOutput()
	lgr.SetOutput(output)

	if len(processors) > 0 {
		var p Processor = func(fields map[string]interface{}, ctx context.Context) *logrus.Entry {
			return lgr.Logger.WithFields(fields)
		}
		for i := len(processors); i > 0; i-- {
			p = processors[i-1].Process(p)
		}
		lgr.processor = p
	}

	return lgr, func() { _ = output.Close() }
}

// GetRequestCtx may be used for extract context values.
func (l *LogrusLogger) GetRequestCtx() context.Context {
	return l.requestCtx
}

// SetRequestCtx must be called on the server while a request processing start.
func (l *LogrusLogger) SetRequestCtx(ctx context.Context) {
	l.requestCtx = ctx
}

func (l *LogrusLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	if l.processor != nil {
		return l.processor(fields, l.requestCtx)
	}
	return l.Logger.WithFields(fields)
}
