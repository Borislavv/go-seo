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
	ctx context.Context
}

func NewInitCtxMiddleware(ctx context.Context) *InitCtxMiddleware {
	return &InitCtxMiddleware{ctx: ctx}
}

func (m *InitCtxMiddleware) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		reqCtx := context.WithValue(m.ctx, values.RequestStartedAt, time.Now())

		ctx.SetUserValue(values.CtxKey, reqCtx)

		next(ctx)
	}
}

/* ------------------------------------------------------------------------------------------------------------------ */

type EnrichCtxByIDMiddleware struct {
	ctx    context.Context
	logger logger.Logger
}

func NewEnrichCtxByIDMiddleware(ctx context.Context, logger logger.Logger) *EnrichCtxByIDMiddleware {
	return &EnrichCtxByIDMiddleware{ctx: ctx, logger: logger}
}

func (m *EnrichCtxByIDMiddleware) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		reqID := string(ctx.Request.Header.Peek(values.RequestIDHeader))
		if reqID == "" {
			m.logger.Errorf(m.ctx, "%v is not exists into request headers "+
				"(provided a new one), some part of logs my be lost", values.RequestIDHeader)
			reqID = uuid.V4()
		}

		reqCtx, ok := ctx.UserValue(values.CtxKey).(context.Context)
		if !ok {
			m.logger.Error(m.ctx, "context.Context is not present into fasthttp.RequestCtx "+
				"(provided a default ctx), some part of logs my be lost")
			reqCtx = m.ctx
		}

		reqCtx = context.WithValue(reqCtx, values.RequestIDKey, reqID)

		ctx.SetUserValue(values.CtxKey, reqCtx)

		next(ctx)
	}
}

/* ------------------------------------------------------------------------------------------------------------------ */

type EnrichCtxByGUIDMiddleware struct {
	ctx    context.Context
	logger logger.Logger
}

func NewEnrichLoggerCtxByGUIDMiddleware(ctx context.Context, logger logger.Logger) *EnrichCtxByGUIDMiddleware {
	return &EnrichCtxByGUIDMiddleware{ctx: ctx, logger: logger}
}

func (m *EnrichCtxByGUIDMiddleware) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		reqGUID := string(ctx.Request.Header.Peek(values.RequestGUIDHeader))
		if reqGUID == "" {
			m.logger.Errorf(m.ctx, "%v is not exists into request headers "+
				"(provided a new one), some part of logs my be lost", values.RequestGUIDHeader)
			reqGUID = uuid.V4()
		}

		reqCtx, ok := ctx.UserValue(values.CtxKey).(context.Context)
		if !ok {
			m.logger.Error(m.ctx, "context.Context is not present into fasthttp.RequestCtx "+
				"(provided a default ctx), some part of logs my be lost")
			reqCtx = m.ctx
		}

		reqCtx = context.WithValue(reqCtx, values.RequestGUIDKey, reqGUID)

		ctx.SetUserValue(values.CtxKey, reqCtx)

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
	ctx    context.Context
}

func NewLogRequestMiddleware(ctx context.Context, logger logger.Logger) *LogRequestMiddleware {
	return &LogRequestMiddleware{ctx: ctx, logger: logger}
}

func (m *LogRequestMiddleware) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		reqCtx, ok := ctx.UserValue(values.CtxKey).(context.Context)
		if !ok {
			m.logger.Error(m.ctx, "context.Context is not present into fasthttp.RequestCtx "+
				"(provided a default ctx), some part of logs my be lost")
			reqCtx = m.ctx
		}

		m.logger.WithFields(reqCtx, map[string]interface{}{
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
