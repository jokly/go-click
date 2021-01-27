package service

import (
	"context"
	"fmt"
)

type logService struct{}

func NewLogService() Service {
	return logService{}
}

func (svc logService) Send(_ context.Context, event interface{}) error {
	fmt.Printf("%v\n", event)

	return nil
}
