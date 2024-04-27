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
	return pq[i].Priority.Before(pq[j].Priority)
}

func (pq Queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *Queue) Push(x any) {
	value := x.(*QueueItem)
	*pq = append(*pq, value)
}

func (pq *Queue) Pop() any {
	n := pq.Len()
	ret := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]

	return ret.Service
}

func NewQueue() *Queue {
	var queue = &Queue{}

	heap.Init(queue)

	return queue
}
