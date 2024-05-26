package logger

import (
	"context"
	"github.com/Borislavv/go-seo/internal/shared/values"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/savsgio/gotils/uuid"
	"github.com/sirupsen/logrus"
)

/* ------------------------------------------------------------------------------------------------------------------ */

type RequestIDProcessor struct {
}

func NewRequestIDProcessor() *RequestIDProcessor {
	return &RequestIDProcessor{}
}

func (p *RequestIDProcessor) Process(next logger.Processor) logger.Processor {
	return func(fields map[string]interface{}, ctx context.Context) *logrus.Entry {
		requestID := ctx.Value(values.RequestIDKey)
		if requestID == nil {
			requestID = uuid.V4()
		} else {
			requestID = requestID.(string)
		}
		fields[values.RequestIDKey] = requestID

		return next(fields, ctx)
	}
}

/* ------------------------------------------------------------------------------------------------------------------ */

type RequestGUIDProcessor struct {
}

func NewRequestGUIDProcessor() *RequestGUIDProcessor {
	return &RequestGUIDProcessor{}
}

func (p *RequestGUIDProcessor) Process(next logger.Processor) logger.Processor {
	return func(fields map[string]interface{}, ctx context.Context) *logrus.Entry {
		requestGUID := ctx.Value(values.RequestGUIDKey)
		if requestGUID == nil {
			requestGUID = uuid.V4()
		} else {
			requestGUID = requestGUID.(string)
		}
		fields[values.RequestGUIDKey] = requestGUID

		return next(fields, ctx)
	}
}
