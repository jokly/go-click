package adapter

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

const LogAdapterName = "log"

type logAdapter struct {
	logger log.Logger
}

func MakeLogAdapter(logger log.Logger) Adapter {
	return &logAdapter{
		logger: logger,
	}
}

func (svc *logAdapter) Send(event interface{}) error {
	level.Info(svc.logger).Log("event", &event)

	return nil
}
