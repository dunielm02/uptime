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

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Less(i, j int) bool {
	return (*q)[i].Priority.Before((*q)[j].Priority)
}

func (q *Queue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *Queue) Push(x any) {
	value := x.(*QueueItem)
	(*q) = append((*q), value)
}

func (q *Queue) Pop() any {
	n := q.Len()
	ret := (*q)[n-1]
	(*q) = (*q)[0:(n - 1)]
	return ret
}

func NewQueue() *Queue {
	var queue = &Queue{}

	heap.Init(queue)

	return queue
}
