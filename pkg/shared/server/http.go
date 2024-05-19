package server

import (
	"github.com/Borislavv/go-seo/internal/pagedata/infrastructure/api/v1/http/controller"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type HTTP struct {
	logger logger.Logger
}

func NewHTTP(logger logger.Logger) *HTTP {
	return &HTTP{
		logger: logger,
	}
}

func (s *HTTP) ListenAndServe() {
	r := router.New()
	controller.NewPagedataGetController(s.logger).AddRoute(r)
	if err := fasthttp.ListenAndServe(":8085", r.Handler); err != nil {
		logrus.Error(err)
	}
}
