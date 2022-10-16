package pyramid

import (
	"container/heap"
)

type Heap struct {
	queue Queue
}

func (h *Heap) Push(value any) {
	heap.Push(&h.queue, value)
}

func (h *Heap) Pop() any {
	return heap.Pop(&h.queue)
}

func (h *Heap) Length() int {
	return h.queue.Len()
}

func NewHeap() Heap {
	var pq Heap

	heap.Init(&pq.queue)

	return pq
}
