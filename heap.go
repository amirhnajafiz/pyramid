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

// Update
// the queue list.
func (h *Heap[T]) Update(old any, new any, ef equalFunction) {
	index := h.find(old, ef)
	if index == -1 {
		return
	}

	h.queue.update(h.queue.list[index], new.(T))
}

func (h *Heap[T]) find(object any, equalFunction equalFunction) int {
	for index, obj := range h.queue.list {
		if equalFunction(obj.value, object.(T)) {
			return index
		}
	}

	return -1
}

// NewHeap
// creates a new heap of any type.
func NewHeap[T any](compareFunction compareFunction[T]) Heap[T] {
	var pq Heap[T]

	// setting the compare function
	pq.queue.compareFunction = compareFunction

	// registering the
	heap.Init(&pq.queue)

	return pq
}
