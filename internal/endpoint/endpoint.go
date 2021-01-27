package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jokly/go-click/internal/service"
)

type Endpoints struct {
	SendEndpoint endpoint.Endpoint
}

func New(svc service.Service) Endpoints {
	sendEndpoint := MakeSendEndpoint(svc)

	return Endpoints{
		SendEndpoint: sendEndpoint,
	}
}

func MakeSendEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendRequest)
		err := svc.Send(ctx, req)

		return SendResponse{Err: err}, nil
	}
}

type SendRequest interface{}

type SendResponse struct {
	Err error `json:"-"`
}
