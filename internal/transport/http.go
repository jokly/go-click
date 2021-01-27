package transport

import (
	"context"
	"encoding/json"
	"net/http"

	kitendpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/jokly/go-click/internal/endpoint"
)

func MakeHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Handle("/send", httptransport.NewServer(
		endpoints.SendEndpoint,
		decodeHTTPSendRequest,
		encodeHTTPResponse,
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
		errorEncoder(ctx, f.Failed(), w)

		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

type errorWrapper struct {
	Error string `json:"error"`
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(convertErrorToCode(err))
	_ = json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func convertErrorToCode(_ error) int {
	return http.StatusInternalServerError
}
