package service

import (
	"context"
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jokly/go-click/internal/service/adapter"
)

type Service interface {
	Send(ctx context.Context, event interface{}) error
	Stop()
}

type SenderService struct {
	adapter adapter.Adapter
	logger  log.Logger
}

func MakeSenderService(adapter adapter.Adapter, logger log.Logger) Service {
	return &SenderService{
		adapter: adapter,
		logger:  logger,
	}
}

func (s *SenderService) Send(_ context.Context, event interface{}) error {
	return s.adapter.Send(event)
}

func (s *SenderService) Stop() {
	level.Info(s.logger).Log("msg", "Stop sender...")
}

type SenderPoolService struct {
	adapter    adapter.Adapter
	numWorkers uint8
	wg         *sync.WaitGroup
	eventsChan chan interface{}
	logger     log.Logger
}

func MakeSenderPoolServcie(adapter adapter.Adapter, numWorkers uint8, logger log.Logger) Service {
	sender := &SenderPoolService{
		adapter:    adapter,
		numWorkers: numWorkers,
		wg:         &sync.WaitGroup{},
		eventsChan: make(chan interface{}, numWorkers),
		logger:     logger,
	}

	sender.serve()

	return sender
}

func (s *SenderPoolService) Send(_ context.Context, event interface{}) error {
	s.eventsChan <- event

	return nil
}

func (s *SenderPoolService) Stop() {
	level.Info(s.logger).Log("msg", "Stop sender...")

	close(s.eventsChan)
	s.wg.Wait()
}

func (s *SenderPoolService) serve() {
	for i := 0; i < int(s.numWorkers); i++ {
		s.wg.Add(1)

		go s.consume()
	}
}

func (s *SenderPoolService) consume() {
	defer s.wg.Done()

	for event := range s.eventsChan {
		if err := s.adapter.Send(event); err != nil {
			level.Error(s.logger).Log("err", err)
		}
	}
}
