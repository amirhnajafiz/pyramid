package pyramid

import (
	"container/heap"
)

type Heap[T any] struct {
	queue Queue[T]
}

func (h *Heap[T]) Push(value any) {
	heap.Push(&h.queue, value)
}

func (h *Heap[T]) Pop() any {
	return heap.Pop(&h.queue)
}

func (h *Heap[T]) Length() int {
	return h.queue.Len()
}

func NewHeap[T any](compareFunction compareFunction) Heap[T] {
	var pq Heap[T]

	pq.queue.compareFunction = compareFunction

	heap.Init(&pq.queue)

	return pq
}
