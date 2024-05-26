package seo

import (
	"context"
	"github.com/Borislavv/go-seo/internal/pagedata/infrastructure/api/v1/http/controller"
	internallogger "github.com/Borislavv/go-seo/internal/shared/logger"
	internalserver "github.com/Borislavv/go-seo/internal/shared/server"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/Borislavv/go-seo/pkg/shared/server"
	"github.com/Borislavv/go-seo/pkg/shared/shutdown"
	"sync"
)

type App struct {
}

func (s *App) Run() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	lgr, clsLgrFn := logger.NewLogger(ctx, processors())
	defer clsLgrFn()

	server.
		NewHTTP(
			lgr,
			controllers(lgr),
			middlewares(ctx, lgr),
		).
		ListenAndServe(ctx, wg)

	shutdown.
		NewGraceful(cancel).
		ListenAndCancel()

	wg.Wait()
}

// controllers returns a slice of server.HttpController[s] for http server (handlers).
func controllers(lgr logger.Logger) []server.HttpController {
	return []server.HttpController{
		controller.NewPagedataGetController(lgr),
	}
}

// middlewares returns a slice of server.HttpMiddleware[s] which will executes in reverse order before handling request.
func middlewares(ctx context.Context, lgr logger.Logger) []server.HttpMiddleware {
	return []server.HttpMiddleware{
		/** exec 1st. */ internalserver.NewInitCtxMiddleware(ctx, lgr),
		/** exec 2nd. */ internalserver.NewEnrichLoggerCtxByIDMiddleware(ctx, lgr),
		/** exec 3rd. */ internalserver.NewEnrichLoggerCtxByGUIDMiddleware(ctx, lgr),
		/** exec 4th. */ internalserver.NewLogRequestMiddleware(lgr),
		/** exec 5th. */ internalserver.NewApplicationJsonMiddleware(),
	}
}

// processors returns a slice of logger.FieldsProcessor[s] which mutate fields map.
func processors() []logger.FieldsProcessor {
	return []logger.FieldsProcessor{
		/** exec 1st. */ internallogger.NewRequestIDProcessor(),
		/** exec 2nd. */ internallogger.NewRequestGUIDProcessor(),
	}
}
