package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jokly/go-click/internal/service"
)

type Endpoints struct {
	SendEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc service.Service) Endpoints {
	return Endpoints{
		SendEndpoint: makeSendEndpoint(svc),
	}
}

func makeSendEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendRequest)
		err := svc.Send(ctx, req)

		return SendResponse{Error: err}, nil
	}
}
