package pyramid

import "container/heap"

// Queue
// is a list of items.
type Queue[T any] struct {
	list            []*item[T]
	compareFunction compareFunction
}

// Len
// returns the Queue size.
func (pq Queue[T]) Len() int {
	return len(pq.list)
}

// Less
// comparator function.
func (pq Queue[T]) Less(i, j int) bool {
	return pq.compareFunction(pq.list[i], pq.list[j])
}

// Swap
// function for Queue.
func (pq Queue[T]) Swap(i, j int) {
	pq.list[i], pq.list[j] = pq.list[j], pq.list[i]
	pq.list[i].index = i
	pq.list[j].index = j
}

// Push
// new items to Queue list.
func (pq *Queue[T]) Push(x any) {
	n := len(pq.list)

	i := x.(*item[T])
	i.index = n

	pq.list = append(pq.list, i)
}

// Pop
// extract one item from Queue list.
func (pq *Queue[T]) Pop() any {
	old := pq.list
	n := len(old)
	i := old[n-1]

	old[n-1] = nil // avoid memory leak
	i.index = -1   // for safety
	pq.list = old[0 : n-1]

	return i
}

// update
// modifies the priority and value of an Item in the queue.
func (pq *Queue[T]) update(i *item[T], value T) {
	i.value = value

	heap.Fix(pq, i.index)
}
