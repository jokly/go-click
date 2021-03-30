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

func (adap *logAdapter) Send(event interface{}) error {
	level.Info(adap.logger).Log("event", &event)

	return nil
}

func (adap *logAdapter) Stop() error {
	return nil
}
