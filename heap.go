package pyramid

import (
	"container/heap"
)

// Heap
// type which manages the heap list of any type.
type Heap[T any] struct {
	// heap has a queue to manage the items init
	queue Queue[T]
}

// Push
// items to heap.
func (h *Heap[T]) Push(value any) {
	heap.Push(&h.queue, value)
}

// Pop
// items from list.
func (h *Heap[T]) Pop() any {
	return heap.Pop(&h.queue)
}

// Length
// heap length.
func (h *Heap[T]) Length() int {
	return h.queue.Len()
}

// NewHeap
// creates a new heap of any type.
func NewHeap[T any](compareFunction compareFunction) Heap[T] {
	var pq Heap[T]

	// setting the compare function
	pq.queue.compareFunction = compareFunction

	// registering the
	heap.Init(&pq.queue)

	return pq
}
