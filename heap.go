package main

import "container/heap"

type Heap []*Item

func (pq Heap) Len() int {
	return len(pq)
}

func (pq Heap) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq Heap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *Heap) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *Heap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *Heap) update(item *Item, value int) {
	item.value = value
	heap.Fix(pq, item.index)
}
