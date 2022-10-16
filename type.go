package pyramid

// An Item is something we manage in a priority queue.
type Item struct {
	// each item has a value of its own that could be anything.
	value any
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// CompareFunction
// is used to compare items in priority queue.
type CompareFunction func(*Item, *Item) bool
