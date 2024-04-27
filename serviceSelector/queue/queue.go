package queue

import (
	"container/heap"
	"lifeChecker/checkers"
	"time"
)

type QueueItem struct {
	Priority time.Time
	Service  checkers.LifeChecker
}

type Queue []*QueueItem

func (pq Queue) Len() int {
	return len(pq)
}

func (pq Queue) Less(i, j int) bool {
	return pq[i].Priority.Compare(pq[i].Priority) > 0
}

func (pq Queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *Queue) Push(x any) {
	item := x.(*QueueItem)
	*pq = append(*pq, item)
}

func (pq *Queue) Pop() any {
	old := *pq
	n := len(old)
	if n == 0 {
		return nil
	}
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item.Service
}

func NewQueue() Queue {
	var queue = Queue{}

	heap.Init(&queue)

	return Queue{}
}
