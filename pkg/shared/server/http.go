package server

import (
	"context"
	"errors"
	"github.com/Borislavv/go-seo/internal/pagedata/infrastructure/api/v1/http/controller"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

type HTTP struct {
	logger logger.Logger
}

func NewHTTP(logger logger.Logger) *HTTP {
	return &HTTP{
		logger: logger,
	}
}

func (s *HTTP) ListenAndServe(ctx context.Context, wg *sync.WaitGroup) {
	r := router.New()
	controller.NewPagedataGetController(s.logger).AddRoute(r)
	server := fasthttp.Server{Handler: r.Handler}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(":8087"); err != nil {
			s.logger.Error(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			s.logger.Info("http server was stopped")
			wg.Done()
		}()

		<-ctx.Done()

		sctx, cancel := context.WithTimeout(ctx, time.Second*15)
		defer cancel()

		if err := server.ShutdownWithContext(sctx); err != nil {
			if errors.Is(err, context.Canceled) {
				s.logger.Error(err)
			}
		}
	}()

	s.logger.Info("http server was started")
}
