package controller

import (
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const PagedataGetPath = "/pagedata"

type PagedataGetController struct {
	logger logger.Logger
}

func NewPagedataGetController(logger logger.Logger) *PagedataGetController {
	return &PagedataGetController{
		logger: logger,
	}
}

func (c *PagedataGetController) Get(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	c.logger.Info("Succeeded response")

	if _, err := ctx.WriteString("Success response"); err != nil {
		c.logger.Error(err)
	}
}

func (c *PagedataGetController) AddRoute(router *router.Router) {
	router.GET(PagedataGetPath, c.Get)
}
