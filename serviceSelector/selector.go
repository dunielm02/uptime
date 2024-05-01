package serviceSelector

import (
	"container/heap"
	"lifeChecker/checkers"
	"lifeChecker/serviceSelector/queue"
	"sync"
	"time"
)

type Selector struct {
	mu   *sync.RWMutex
	list *queue.Queue
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

func (s *Selector) NextItem() (checkers.LifeChecker, time.Time) {
	if s.len() == 0 {
		return nil, time.Time{}
	}

	s.mu.Lock()
	ret := heap.Pop(s.list).(*queue.QueueItem)
	s.mu.Unlock()

	return ret.Service, ret.Priority
}

func (s *Selector) len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.list.Len()
}

// func (s *Selector) RunChecking(database.DB) {
// 	var lifeResult = make(chan database.TimeSerie)
// 	for {
// 		if s.len() == 0 {
// 			continue
// 		}

// 		service, initTime := s.NextItem()

// 		time.Sleep(time.Until(initTime))

// 		go func(serv checkers.LifeChecker) {
// 			alive := true
// 			requestDuration, err := serv.CheckLife()

// 			if err != nil {
// 				if !serv.IsInverted() {
// 					alive = false
// 				}
// 			}

// 			lifeResult <- database.TimeSerie{
// 				Name:        serv.GetName(),
// 				RequestTime: requestDuration,
// 				Alive:       alive,
// 			}
// 		}(service)
// 	}
// }
