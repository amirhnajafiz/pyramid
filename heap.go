package main

import "container/heap"

// Heap
// is a list of items.
type Heap []*Item

// Len
// returns the heap size.
func (pq Heap) Len() int {
	return len(pq)
}

// Less
// comparator function.
func (pq Heap) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

// Swap
// function for heap.
func (pq Heap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push
// new items to heap list.
func (pq *Heap) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop
// extract one item from heap list.
func (pq *Heap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]

	return item
}

// update modifies the priority and
// value of an Item in the queue.
func (pq *Heap) update(item *Item, value int) {
	item.value = value

	heap.Fix(pq, item.index)
}
