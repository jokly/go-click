package main

import (
	"flag"
	"fmt"
	"net/http"

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

	logAdapter := adapter.MakeLogAdapter(logger)
	logService := service.MakeSenderPoolServcie(logAdapter, 3, logger)
	endpoints := endpoint.MakeEndpoints(logService)
	httpHandler := transport.MakeHTTPHandler(endpoints, logger)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HTTP.Port),
		Handler: httpHandler,
	}

	err = server.ListenAndServe()
	level.Error(logger).Log("err", err)
}
