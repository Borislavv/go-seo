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

func NewHTTP(
	logger logger.Logger,
	controllers []HttpController,
	middlewares []HttpMiddleware,
) *HTTP {
	s := &HTTP{logger: logger, config: new(Config).Load()}
	s.initServer(s.buildRouter(controllers), middlewares)
	return s
}

func (s *HTTP) ListenAndServe(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go s.serve(wg)

	wg.Add(1)
	go s.shutdown(wg, ctx)
}

func (s *HTTP) serve(wg *sync.WaitGroup) {
	defer wg.Done()

	s.logger.Infof("http server was started on port http://0.0.0.0%v", s.config.Port)

	if err := s.server.ListenAndServe(s.config.Port); err != nil {
		s.logger.Error(err)
	}
}

func (s *HTTP) shutdown(wg *sync.WaitGroup, ctx context.Context) {
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
}

func (s *HTTP) buildRouter(controllers []HttpController) *router.Router {
	r := router.New()
	for _, controller := range controllers {
		controller.AddRoute(r)
	}
	return r
}

func (s *HTTP) initServer(r *router.Router, mdws []HttpMiddleware) {
	h := r.Handler

	for i := len(mdws) - 1; i >= 0; i-- {
		h = mdws[i].Middleware(h)
	}

	s.server = &fasthttp.Server{Handler: h}
}
