package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type logService struct {
	logger log.Logger
}

func NewLogService(logger log.Logger) Service {
	return logService{
		logger: logger,
	}
}

func (svc logService) Send(_ context.Context, event interface{}) error {
	level.Info(svc.logger).Log("event", event)

	return nil
}
