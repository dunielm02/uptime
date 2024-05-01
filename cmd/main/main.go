package main

import (
	"container/heap"
	"fmt"
	que "lifeChecker/serviceSelector/queue"
	"math/rand"
	"sync"
	"time"
)

type mock_checkLifeService struct {
	name string
}

func (s *mock_checkLifeService) GetName() string {
	return s.name
}

func (s *mock_checkLifeService) CheckLife() (time.Duration, error) { return time.Duration(0), nil }
func (s *mock_checkLifeService) IsInverted() bool                  { return true }
func (s *mock_checkLifeService) GetQueueTime() time.Duration       { return time.Duration(0) }

func main() {
	var queue = que.NewQueue()
	wg := sync.WaitGroup{}

	for range 10 {
		value := time.Now().Add(time.Duration(rand.Int()))
		wg.Add(1)
		go func() {
			defer wg.Done()
			heap.Push(queue, &que.QueueItem{
				Priority: value,
				Service: &mock_checkLifeService{
					name: fmt.Sprint(value.Unix()),
				},
			})
		}()
	}

	wg.Wait()

	fmt.Println("=======================")

	wg2 := sync.WaitGroup{}

	for range 10 {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			heap.Pop(queue)
		}()
	}

	wg2.Wait()
}
