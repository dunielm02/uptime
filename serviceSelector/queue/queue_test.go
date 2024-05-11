package queue

import (
	"container/heap"
	"fmt"
	"lifeChecker/tests/mocks"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("testing that the Queues is Ordered", func(t *testing.T) {
		var queue = NewQueue()
		var slice = []time.Time{}
		for range 10 {
			value := time.Now().Add(time.Duration(rand.Int()))
			slice = append(slice, value)
			heap.Push(queue, &QueueItem{
				Priority: value,
				Service: &mocks.Mock_checkLifeService{
					Name: fmt.Sprint(value.Unix()),
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
	})
}
