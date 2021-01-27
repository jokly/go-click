package service

import "context"

type Service interface {
	Send(ctx context.Context, event interface{}) error
}
