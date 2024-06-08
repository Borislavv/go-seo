package controller

import (
	"context"
	"encoding/json"
	"github.com/Borislavv/go-seo/internal/shared/infrastructure/values"
	"github.com/Borislavv/go-seo/pkg/shared/api/http/controller"
	"github.com/Borislavv/go-seo/pkg/shared/cache"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"time"
)

const PagedataGetPath = "/pagedata"

type PagedataGet struct {
	abstract *controller.Abstract
	ctx      context.Context
	cache    cache.Cacher
	logger   logger.Logger
}

func NewPagedataGet(ctx context.Context, cache cache.Cacher, logger logger.Logger) *PagedataGet {
	return &PagedataGet{
		abstract: controller.NewAbstractController(logger),
		ctx:      ctx,
		cache:    cache,
		logger:   logger,
	}
}

func (c *PagedataGet) Get(ctx *fasthttp.RequestCtx) {
	reqCtx, ok := ctx.UserValue(values.CtxKey).(context.Context)
	if !ok {
		c.logger.Error(c.ctx, "context.Context is not exists into the fasthttp.RequestCtx "+
			"(provided default ctx), some part of logs may be lost")
		reqCtx = c.ctx
	}

	data, err := c.cache.Get(
		string(ctx.Request.Header.Method())+" "+ctx.Request.URI().String(),
		func(item cache.CacheItem) (data interface{}, err error) {
			item.SetTTL(time.Minute * 30)
			return c.get()
		},
	)
	if err != nil {
		c.logger.Error(c.ctx, "failed to get the data from the cache: "+err.Error())
		c.abstract.InternalServerError(ctx)
		return
	}

	byteSlice, ok := data.([]byte)
	if !ok {
		c.logger.Error(c.ctx, "failed to convert the data to byte slice")
		c.abstract.InternalServerError(ctx)
		return
	}

	if _, err = ctx.Write(byteSlice); err != nil {
		c.logger.Errorf(reqCtx, "ctx.Write fail: %v", err.Error())
		c.abstract.InternalServerError(ctx)
	} else {
		ctx.SetStatusCode(fasthttp.StatusOK)
		c.logger.Info(reqCtx, "http response")
	}
}

func (c *PagedataGet) AddRoute(router *router.Router) {
	router.GET(PagedataGetPath, c.Get)
}

func (c *PagedataGet) get() ([]byte, error) {
	datum := make(map[string]map[string]interface{}, 1)
	datum["data"] = make(map[string]interface{}, 1)
	datum["data"]["success"] = true
	return json.Marshal(datum)
}
