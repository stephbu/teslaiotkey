package logging

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type loggerKeyType int

const loggerKey loggerKeyType = iota

var logger *logrus.Logger

func init() {
	logger = logrus.StandardLogger()

}

func NewContext(ctx context.Context, fields logrus.Fields) context.Context {
	return context.WithValue(ctx, loggerKey, WithContext(ctx).WithFields(fields))
}

func WithContext(ctx context.Context) *logrus.Logger {
	if ctx == nil {
		return logger
	}

	if ctxLogger, ok := ctx.Value(loggerKey).(*logrus.Logger); ok {
		return ctxLogger
	} else {
		return logger
	}
}

func CreateRequestContext(ctx context.Context) context.Context {
	return NewContext(ctx, logrus.Fields{"requestID": uuid.New()})
}
