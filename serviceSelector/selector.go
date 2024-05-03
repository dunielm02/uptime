package serviceSelector

import (
	"container/heap"
	"context"
	"lifeChecker/checkers"
	"lifeChecker/config"
	"lifeChecker/database"
	"lifeChecker/serviceSelector/queue"
	"sync"
	"time"
)

type Selector struct {
	mu   *sync.RWMutex
	list *queue.Queue
}

func SelectorFromConfig(cfg []config.ServiceConfig) *Selector {
	selector := NewSelector()

	for _, i := range cfg {
		selector.Insert(checkers.GetFromConfig(i))
	}

	return selector
}

func NewSelector() *Selector {
	return &Selector{
		mu:   &sync.RWMutex{},
		list: queue.NewQueue(),
	}
}

func (s *Selector) Insert(service checkers.LifeChecker) {
	var item = queue.QueueItem{
		Priority: time.Now().Add(service.GetQueueTime()),
		Service:  service,
	}
	s.mu.Lock()
	heap.Push(s.list, &item)
	s.mu.Unlock()
}

func (s *Selector) NextItem() checkers.LifeChecker {
	s.mu.Lock()
	ret := heap.Pop(s.list).(*queue.QueueItem)
	s.mu.Unlock()

	return ret.Service
}

func (s *Selector) ItsTimeToCheck() bool {
	if s.len() == 0 {
		return false
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	return time.Now().Compare((*s.list)[0].Priority) >= 0
}

func (s *Selector) len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.list.Len()
}

func (s *Selector) RunChecking(ctx context.Context, db database.DB) {
	var lifeResult = make(chan database.TimeSerie)
	go db.WriteTimeSerieFromChannel(ctx, lifeResult)
	for ctx.Err() == nil {
		for !s.ItsTimeToCheck() {
		}

		service := s.NextItem()

		go func(serv checkers.LifeChecker) {
			alive := true
			requestDuration, err := serv.CheckLife()

			if err != nil {
				if !serv.IsInverted() {
					alive = false
				}
			}

			toSend := database.TimeSerie{
				Name:        serv.GetName(),
				RequestTime: requestDuration,
				Alive:       alive,
			}

			lifeResult <- toSend

			s.Insert(serv)
		}(service)
	}
}
