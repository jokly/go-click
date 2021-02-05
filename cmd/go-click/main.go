package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jokly/go-click/internal/endpoint"
	"github.com/jokly/go-click/internal/service"
	"github.com/jokly/go-click/internal/service/adapter"
	"github.com/jokly/go-click/internal/transport"
	"github.com/jokly/go-click/internal/util"
)

func main() {
	configFilePath := flag.String("config", "", "path to config file (e.g. /config/config.yaml)")
	flag.Parse()

	config, err := loadConfig(*configFilePath)
	if err != nil {
		panic(err)
	}

	zLogger, logger := util.InitLogger(config.Logger.MinLevel)

	defer func() {
		_ = zLogger.Sync()
	}()

	svc := makeService(config, logger)
	endpoints := endpoint.MakeEndpoints(svc)
	httpHandler := transport.MakeHTTPHandler(endpoints, logger)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HTTP.Port),
		Handler: httpHandler,
	}

	err = server.ListenAndServe()
	level.Error(logger).Log("err", err)
}

func makeService(config *Config, logger log.Logger) service.Service {
	var adap adapter.Adapter

	switch config.Sender.Adapter {
	case adapter.LogAdapterName:
		adap = adapter.MakeLogAdapter(logger)
	default:
		adap = adapter.MakeLogAdapter(logger)
	}

	var svc service.Service
	if config.Sender.IsPool {
		svc = service.MakeSenderPoolServcie(adap, config.Sender.NumWorkers, logger)
	} else {
		svc = service.MakeSenderService(adap, logger)
	}

	return svc
}
