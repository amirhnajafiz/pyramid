package pyramid

import "container/heap"

// Queue
// is a list of items.
type Queue struct {
	list            []*Item
	compareFunction func(*Item, *Item) bool
}

// Len
// returns the Queue size.
func (pq Queue) Len() int {
	return len(pq.list)
}

// Less
// comparator function.
func (pq Queue) Less(i, j int) bool {
	return pq.compareFunction(pq.list[i], pq.list[j])
}

// Swap
// function for Queue.
func (pq Queue) Swap(i, j int) {
	pq.list[i], pq.list[j] = pq.list[j], pq.list[i]
	pq.list[i].index = i
	pq.list[j].index = j
}

// Push
// new items to Queue list.
func (pq *Queue) Push(x any) {
	n := len(pq.list)

	item := x.(*Item)
	item.index = n

	pq.list = append(pq.list, item)
}

// Pop
// extract one item from Queue list.
func (pq *Queue) Pop() any {
	old := pq.list
	n := len(old)
	item := old[n-1]

	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.list = old[0 : n-1]

	return item
}

// update modifies the priority and
// value of an Item in the queue.
func (pq *Queue) update(item *Item, value int) {
	item.value = value

	heap.Fix(pq, item.index)
}
