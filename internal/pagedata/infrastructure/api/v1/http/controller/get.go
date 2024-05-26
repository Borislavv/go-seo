package controller

import (
	"context"
	"encoding/json"
	"github.com/Borislavv/go-seo/internal/shared/values"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const PagedataGetPath = "/pagedata"

type PagedataGetController struct {
	ctx    context.Context
	logger logger.Logger
}

func NewPagedataGetController(ctx context.Context, logger logger.Logger) *PagedataGetController {
	return &PagedataGetController{
		ctx:    ctx,
		logger: logger,
	}
}

func (c *PagedataGetController) Get(ctx *fasthttp.RequestCtx) {
	reqCtx, ok := ctx.UserValue(values.CtxKey).(context.Context)
	if !ok {
		c.logger.Error(c.ctx, "context.Context is not exists into the fasthttp.RequestCtx "+
			"(provided default ctx), some part of logs may be lost")
		reqCtx = c.ctx
	}

	data := make(map[string]map[string]bool, 1)
	data["data"] = make(map[string]bool, 1)
	data["data"]["success"] = true

	b, err := json.Marshal(data)
	if err != nil {
		c.logger.Errorf(c.ctx, "json marshal fail: %v", err.Error())
	}

	if _, err = ctx.Write(b); err != nil {
		c.logger.Errorf(reqCtx, "ctx.Write fail: %v", err.Error())
	} else {
		ctx.SetStatusCode(fasthttp.StatusOK)
		c.logger.Info(reqCtx, "http response")
	}
}

func (c *PagedataGetController) AddRoute(router *router.Router) {
	router.GET(PagedataGetPath, c.Get)
}
