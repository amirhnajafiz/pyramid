package pyramid

import "container/heap"

// Heap
// is a list of items.
type Heap struct {
	list            []*Item
	compareFunction func(*Item, *Item) bool
}

// Len
// returns the heap size.
func (pq Heap) Len() int {
	return len(pq.list)
}

// Less
// comparator function.
func (pq Heap) Less(i, j int) bool {
	return pq.compareFunction(pq.list[i], pq.list[j])
}

// Swap
// function for heap.
func (pq Heap) Swap(i, j int) {
	pq.list[i], pq.list[j] = pq.list[j], pq.list[i]
	pq.list[i].index = i
	pq.list[j].index = j
}

// Push
// new items to heap list.
func (pq *Heap) Push(x any) {
	n := len(pq.list)

	item := x.(*Item)
	item.index = n

	pq.list = append(pq.list, item)
}

// Pop
// extract one item from heap list.
func (pq *Heap) Pop() any {
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
func (pq *Heap) update(item *Item, value int) {
	item.value = value

	heap.Fix(pq, item.index)
}
