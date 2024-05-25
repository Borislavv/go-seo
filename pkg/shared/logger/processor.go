package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type FieldsProcessor interface {
	Process(p Processor, ctx context.Context) Processor
}

type Processor func(fields map[string]interface{}) *logrus.Entry
