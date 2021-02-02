package transport

import (
	"context"
	"encoding/json"
	"net/http"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/jokly/go-click/internal/endpoint"
)

func MakeHTTPHandler(endpoints endpoint.Endpoints, logger log.Logger) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	r := mux.NewRouter()

	r.Handle("/send", kithttp.NewServer(
		endpoints.SendEndpoint,
		decodeHTTPSendRequest,
		encodeHTTPResponse,
		options...,
	))

	return r
}

func decodeHTTPSendRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.SendRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	return req, err
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitendpoint.Failer); ok && f.Failed() != nil {
		encodeError(ctx, f.Failed(), w)

		return nil
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(convertErrorToCode(err))
}

func convertErrorToCode(_ error) int {
	return http.StatusInternalServerError
}
