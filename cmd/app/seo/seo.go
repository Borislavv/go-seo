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
			middlewares(lgr),
		).
		ListenAndServe(ctx, wg)

	shutdown.
		NewGraceful(cancel).
		ListenAndCancel()

	wg.Wait()
}

// controllers: returns a slice of server.HttpController[s] for http server (handlers).
func controllers(lgr logger.Logger) []server.HttpController {
	return []server.HttpController{
		controller.NewPagedataGetController(lgr),
	}
}

// middlewares: returns a slice of server.HttpMiddleware[s] which will executes in reverse order before handling request.
func middlewares(lgr logger.Logger) []server.HttpMiddleware {
	return []server.HttpMiddleware{
		internalserver.NewApplicationJsonMiddleware(),
		internalserver.NewLogRequestMiddleware(lgr),
	}
}

// processors: returns a slice of logger.FieldsProcessor[s] which mutate fields map.
func processors() []logger.FieldsProcessor {
	return []logger.FieldsProcessor{
		internallogger.NewRequestIDProcessor(),
		internallogger.NewRequestGUIDProcessor(),
	}
}
