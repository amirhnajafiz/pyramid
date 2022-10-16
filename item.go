package pyramid

// An Item is something we manage in a priority queue.
type Item struct {
	value any
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}
