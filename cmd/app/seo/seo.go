package seo

import (
	"context"
	"github.com/Borislavv/go-seo/internal/pagedata/infrastructure/api/v1/http/controller"
	internalserver "github.com/Borislavv/go-seo/internal/shared/server"
	"github.com/Borislavv/go-seo/internal/shared/values"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/Borislavv/go-seo/pkg/shared/server"
	"github.com/Borislavv/go-seo/pkg/shared/shutdown"
	"sync"
)

type App struct {
}

// Run starts the SEO application.
func (s *App) Run() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	lgr, clsLgrFn := logger.NewLogger(values.ExtraFields)
	defer clsLgrFn()

	server.
		NewHTTP(
			ctx,
			lgr,
			controllers(ctx, lgr),
			middlewares(ctx, lgr),
		).
		ListenAndServe(wg)

	shutdown.
		NewGraceful(cancel).
		ListenAndCancel()

	wg.Wait()
}

// controllers returns a slice of server.HttpController[s] for http server (handlers).
func controllers(ctx context.Context, lgr logger.Logger) []server.HttpController {
	return []server.HttpController{
		controller.NewPagedataGetController(ctx, lgr),
	}
}

// middlewares returns a slice of server.HttpMiddleware[s] which will executes in reverse order before handling request.
func middlewares(ctx context.Context, lgr logger.Logger) []server.HttpMiddleware {
	return []server.HttpMiddleware{
		/** exec 1st. */ internalserver.NewInitCtxMiddleware(ctx),
		/** exec 2nd. */ internalserver.NewEnrichCtxByIDMiddleware(ctx, lgr),
		/** exec 3rd. */ internalserver.NewEnrichLoggerCtxByGUIDMiddleware(ctx, lgr),
		/** exec 4th. */ internalserver.NewLogRequestMiddleware(ctx, lgr),
		/** exec 5th. */ internalserver.NewApplicationJsonMiddleware(),
	}
}
