package server

import (
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/valyala/fasthttp"
)

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
		})
		next(ctx)
	}
}

/* ------------------------------------------------------------------------------------------------------------------ */
