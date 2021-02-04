package service

import (
	"context"

	"github.com/jokly/go-click/internal/service/adapter"
)

type Service interface {
	Send(ctx context.Context, event interface{}) error
}

type SenderService struct {
	adapter adapter.Adapter
}

func MakeSenderService(adapter adapter.Adapter) Service {
	return SenderService{
		adapter: adapter,
	}
}

func (s SenderService) Send(_ context.Context, event interface{}) error {
	return s.adapter.Send(event)
}
