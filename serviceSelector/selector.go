package serviceSelector

import (
	"container/heap"
	"lifeChecker/checkers"
	"lifeChecker/database"
	"lifeChecker/serviceSelector/queue"
	"time"
)

// This should support concurrent access

type Selector struct {
	list *queue.Queue
}

func NewSelector() *Selector {
	return &Selector{
		list: queue.NewQueue(),
	}
}

func (s *Selector) Insert(service checkers.LifeChecker) {
	var item = queue.QueueItem{
		Priority: time.Now().Add(service.GetQueueTime()),
		Service:  service,
	}

	heap.Push(s.list, &item)
}

func (s *Selector) NextItem() (checkers.LifeChecker, time.Time) {
	if s.list.Len() == 0 {
		return nil, time.Time{}
	}

	ret := heap.Pop(s.list).(*queue.QueueItem)

	return ret.Service, ret.Priority
}

func (s *Selector) len() int {
	return len(*s.list)
}

func (s *Selector) RunChecking(database.DB) {
	var lifeResult = make(chan database.TimeSerie)
	for {
		if s.len() == 0 {
			continue
		}

		service, initTime := s.NextItem()

		time.Sleep(time.Until(initTime))

		go func(serv checkers.LifeChecker) {
			alive := true
			requestDuration, err := serv.CheckLife()

			if err != nil {
				if !serv.IsInverted() {
					alive = false
				}
			}

			lifeResult <- database.TimeSerie{
				Name:        serv.GetName(),
				RequestTime: requestDuration,
				Alive:       alive,
			}
		}(service)
	}
}
