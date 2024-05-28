package controller

import (
	"context"
	"github.com/Borislavv/go-seo/internal/shared/values"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/valyala/fasthttp"
)

type Abstract struct {
	ctx    context.Context
	logger logger.Logger
}

func NewAbstractController(logger logger.Logger) *Abstract {
	return &Abstract{logger: logger}
}

func (c *Abstract) InternalServerError(ctx *fasthttp.RequestCtx) {
	if _, err := ctx.WriteString(values.InternalServerError); err != nil {
		c.logger.Error(c.ctx, "failed to write response into the fasthttp.RequestCtx due to "+err.Error())
		return
	}
}
