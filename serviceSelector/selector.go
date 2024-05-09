package serviceSelector

import (
	"container/heap"
	"context"
	"lifeChecker/checkers"
	"lifeChecker/config"
	"lifeChecker/database"
	"lifeChecker/notifications"
	"lifeChecker/serviceSelector/queue"
	"log"
	"sync"
	"time"
)

type Selector struct {
	mu     *sync.RWMutex
	alerts sync.Map
	list   *queue.Queue
}

func SelectorFromConfig(cfg config.Config) *Selector {
	selector := NewSelector()

	for _, i := range cfg.NotificationChannels {
		val, err := notifications.GetNotificationChannelFromConfig(i)
		if err != nil {
			log.Printf("there was an error adding a notification channel: %s", err)
			continue
		}
		selector.alerts.Store(i.Name, val)
	}

	for _, i := range cfg.Services {
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
			alive := checkers.StatusAlive
			requestDuration, err := serv.CheckLife()

			if err != nil {
				if !serv.IsInverted() {
					alive = checkers.StatusDead
				}
			}

			if alive != serv.GetState() {
				s.SendNotifications(serv, alive)
			}

			toSend := database.TimeSerie{
				Name:        serv.GetName(),
				RequestTime: requestDuration,
				Alive:       alive == checkers.StatusAlive,
			}

			lifeResult <- toSend

			s.Insert(serv)
		}(service)
	}
}

func (s *Selector) SendNotifications(service checkers.LifeChecker, alive checkers.State) {
	channels := service.GetNotificationChannelsNames()

	for _, chanName := range channels {
		c, ok := s.alerts.Load(chanName)
		if !ok {
			log.Printf("The notification channel: \"%s\" does not exist.\n", chanName)
		}

		channel := c.(notifications.NotificationChannel)

		if alive == checkers.StatusAlive {
			channel.AliveNotification(service.GetName())
		} else {
			channel.DeadNotification(service.GetName())
		}
	}
}
