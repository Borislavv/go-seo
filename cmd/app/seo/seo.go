package seo

import (
	"context"
	"github.com/Borislavv/go-seo/internal/pagedata/infrastructure/api/v1/http/controller"
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

	loggerService, closeLoggerFunc := logger.NewLogger()
	defer closeLoggerFunc()

	server.
		NewHTTP(loggerService, initControllers(loggerService)).
		ListenAndServe(ctx, wg)

	shutdown.
		NewGraceful(cancel).
		ListenAndCancel()

	wg.Wait()
}

func initControllers(loggerService logger.Logger) []server.HttpController {
	return []server.HttpController{
		controller.NewPagedataGetController(loggerService),
	}
}
