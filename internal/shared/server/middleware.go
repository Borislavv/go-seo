package server

import (
	"context"
	"github.com/Borislavv/go-seo/internal/shared/values"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/savsgio/gotils/uuid"
	"github.com/valyala/fasthttp"
	"time"
)

/* ------------------------------------------------------------------------------------------------------------------ */

type InitCtxMiddleware struct {
	ctx    context.Context
	logger logger.Logger
}

func NewInitCtxMiddleware(ctx context.Context, logger logger.Logger) *InitCtxMiddleware {
	return &InitCtxMiddleware{ctx: ctx, logger: logger}
}

func (m *InitCtxMiddleware) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		reqCtx := context.WithValue(m.ctx, values.RequestStartedAt, time.Now())

		m.logger.SetRequestCtx(reqCtx)

		next(ctx)
	}
}

/* ------------------------------------------------------------------------------------------------------------------ */

type EnrichLoggerCtxByIDMiddleware struct {
	ctx    context.Context
	logger logger.Logger
}

func NewEnrichLoggerCtxByIDMiddleware(ctx context.Context, logger logger.Logger) *EnrichLoggerCtxByIDMiddleware {
	return &EnrichLoggerCtxByIDMiddleware{ctx: ctx, logger: logger}
}

func (m *EnrichLoggerCtxByIDMiddleware) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		requestID := string(ctx.Request.Header.Peek(values.RequestIDHeader))
		if requestID == "" {
			requestID = uuid.V4()
		}

		reqCtx := m.logger.GetRequestCtx()
		reqCtx = context.WithValue(reqCtx, values.RequestIDKey, requestID)

		m.logger.SetRequestCtx(reqCtx)

		next(ctx)
	}
}

/* ------------------------------------------------------------------------------------------------------------------ */

type EnrichLoggerCtxByGUIDMiddleware struct {
	ctx    context.Context
	logger logger.Logger
}

func NewEnrichLoggerCtxByGUIDMiddleware(ctx context.Context, logger logger.Logger) *EnrichLoggerCtxByGUIDMiddleware {
	return &EnrichLoggerCtxByGUIDMiddleware{ctx: ctx, logger: logger}
}

func (m *EnrichLoggerCtxByGUIDMiddleware) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		requestGUID := string(ctx.Request.Header.Peek(values.RequestGUIDHeader))
		if requestGUID == "" {
			requestGUID = uuid.V4()
		}

		reqCtx := m.logger.GetRequestCtx()
		reqCtx = context.WithValue(reqCtx, values.RequestGUIDKey, requestGUID)

		m.logger.SetRequestCtx(reqCtx)

		next(ctx)
	}
}

/* ------------------------------------------------------------------------------------------------------------------ */

type ApplicationJsonMiddleware struct{}

func NewApplicationJsonMiddleware() ApplicationJsonMiddleware {
	return ApplicationJsonMiddleware{}
}

func (ApplicationJsonMiddleware) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")

		next(ctx)
	}
}

/* ------------------------------------------------------------------------------------------------------------------ */

type LogRequestMiddleware struct {
	logger logger.Logger
}

func NewLogRequestMiddleware(logger logger.Logger) *LogRequestMiddleware {
	return &LogRequestMiddleware{logger: logger}
}

func (m *LogRequestMiddleware) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		m.logger.WithFields(map[string]interface{}{
			"method":   string(ctx.Method()),
			"path":     string(ctx.Path()),
			"host":     string(ctx.Host()),
			"remote":   ctx.RemoteIP().String(),
			"connID":   ctx.ConnID(),
			"connTime": ctx.ConnTime().String(),
		}).Info("http request")

		next(ctx)
	}
}

/* ------------------------------------------------------------------------------------------------------------------ */
