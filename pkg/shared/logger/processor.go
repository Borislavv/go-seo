package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type FieldsProcessor interface {
	Process(p Processor) Processor
}

type Processor func(fields map[string]interface{}, ctx context.Context) *logrus.Entry
