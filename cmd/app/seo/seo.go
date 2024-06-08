package seo

import (
	"context"
	pagedatacontroller "github.com/Borislavv/go-seo/internal/pagedata/infrastructure/api/v1/http/controller"
	internalserver "github.com/Borislavv/go-seo/internal/shared/infrastructure/server"
	"github.com/Borislavv/go-seo/internal/shared/infrastructure/values"
	titlecontroller "github.com/Borislavv/go-seo/internal/title/infrastructure/api/v1/http/controller"
	"github.com/Borislavv/go-seo/pkg/shared/cache"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/Borislavv/go-seo/pkg/shared/server"
	"github.com/Borislavv/go-seo/pkg/shared/shutdown"
	"sync"
	"time"
)

type App struct {
}

// Run starts the SEO application.
func (s *App) Run() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	lgr, clsLgrFn := logger.NewLogger(values.ExtraFields)
	defer clsLgrFn()

	chr := cache.NewCache(
		cache.NewMapCacheStorage(ctx),
		cache.NewCacheDisplacer(ctx, time.Second*15),
	)

	server.
		NewHTTP(
			ctx,
			lgr,
			controllers(ctx, chr, lgr),
			middlewares(ctx, lgr),
		).
		ListenAndServe(wg)

	shutdown.
		NewGraceful(cancel).
		ListenAndCancel()

	wg.Wait()
}

// controllers returns a slice of server.HttpController[s] for http server (handlers).
func controllers(ctx context.Context, cache cache.Cacher, logger logger.Logger) []server.HttpController {
	return []server.HttpController{
		pagedatacontroller.NewPagedataGet(ctx, cache, logger),
		titlecontroller.NewTitleGet(ctx, cache, logger),
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
