package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jokly/go-click/internal/endpoint"
	"github.com/jokly/go-click/internal/service"
	"github.com/jokly/go-click/internal/transport"
)

func main() {
	configFilePath := flag.String("config", "", "path to config file (e.g. /config/config.yaml)")
	flag.Parse()

	config, _ := loadConfig(*configFilePath)

	logService := service.NewLogService()
	endpoints := endpoint.MakeEndpoints(logService)
	httpHandler := transport.MakeHTTPHandler(endpoints)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Http.Port),
		Handler: httpHandler,
	}

	_ = server.ListenAndServe()
}
