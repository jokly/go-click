package main

import (
	"net/http"

	"github.com/jokly/go-click/internal/endpoint"
	"github.com/jokly/go-click/internal/service"
	"github.com/jokly/go-click/internal/transport"
)

func main() {
	logService := service.NewLogService()
	endpoints := endpoint.New(logService)
	httpHandler := transport.MakeHTTPHandler(endpoints)

	server := &http.Server{
		Addr:    ":8888",
		Handler: httpHandler,
	}

	server.ListenAndServe()
}
