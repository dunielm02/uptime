package queue

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestQueue(t *testing.T) {
	var queue = NewQueue()
	var slice = []time.Time{}

	for range 10 {
		value := time.Now().Add(time.Duration(rand.Int()))
		slice = append(slice, value)

		heap.Push(queue, &QueueItem{
			Priority: value,
			Service: &mock_checkLifeService{
				name: fmt.Sprint(value.Unix()),
			},
		})
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Before(slice[j])
	})

	for i := range slice {
		queueValue := heap.Pop(queue).(*QueueItem)
		assert.Equal(t, slice[i].Compare(queueValue.Priority), 0)
	}
}
