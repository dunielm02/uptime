package serviceSelector

import (
	"container/heap"
	"lifeChecker/checkers"
	"lifeChecker/serviceSelector/queue"
	"sync"
	"time"
)

type Selector struct {
	mu   *sync.Mutex
	list queue.Queue
}

func NewSelector() Selector {
	return Selector{
		mu:   &sync.Mutex{},
		list: queue.NewQueue(),
	}
}

func (s *Selector) Insert(service checkers.LifeChecker) {
	var item = queue.QueueItem{
		Priority: time.Now().Add(service.GetQueueTime()),
		Service:  service,
	}

	s.mu.Lock()
	heap.Push(&s.list, &item)
	s.mu.Unlock()
}

func (s *Selector) NextItem() checkers.LifeChecker {
	s.mu.Lock()

	ret := heap.Pop(&s.list).(checkers.LifeChecker)

	s.mu.Unlock()

	return ret
}
