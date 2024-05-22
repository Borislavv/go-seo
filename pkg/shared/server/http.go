package server

import (
	"context"
	"errors"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

type HTTP struct {
	server *fasthttp.Server
	config *Config
	logger logger.Logger
}

func NewHTTP(logger logger.Logger, controllers []HttpController) *HTTP {
	s := &HTTP{logger: logger, config: new(Config).Load()}
	s.init(s.buildRouter(controllers))
	return s
}

func (s *HTTP) ListenAndServe(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.server.ListenAndServe(s.config.Port); err != nil {
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

		sctx, cancel := context.WithTimeout(ctx, time.Duration(s.config.ShutDownTimeoutSeconds))
		defer cancel()

		if err := s.server.ShutdownWithContext(sctx); err != nil {
			if errors.Is(err, context.Canceled) {
				s.logger.Info(err)
			} else {
				s.logger.Error(err)
			}
		}
	}()

	s.logger.Infof("http server was started on port http://0.0.0.0%v", s.config.Port)
}

func (s *HTTP) buildRouter(controllers []HttpController) *router.Router {
	r := router.New()
	for _, controller := range controllers {
		controller.AddRoute(r)
	}
	return r
}

func (s *HTTP) init(r *router.Router) {
	s.server = &fasthttp.Server{Handler: r.Handler}
}
