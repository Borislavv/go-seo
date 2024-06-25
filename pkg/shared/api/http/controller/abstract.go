package controller

import (
	"context"
	"github.com/Borislavv/go-seo/internal/shared/infrastructure/values"
	"github.com/Borislavv/go-seo/pkg/shared/api/http/response"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
)

type Abstract struct {
	ctx    context.Context
	logger logger.Logger
}

func NewAbstractController(logger logger.Logger) *Abstract {
	return &Abstract{logger: logger}
}

func (c *Abstract) InternalServerError(w response.WriterOrStringWriter) {
	if _, err := w.WriteString(values.InternalServerError); err != nil {
		c.logger.Error(c.ctx, "failed to write response due to "+err.Error())
		return
	}
}
