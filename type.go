package pyramid

// An Item is something we manage in a priority queue.
type item[T any] struct {
	// each item has a value of its own that could be anything.
	value T
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// CompareFunction
// is used to compare items in priority queue.
type compareFunction func(any, any) bool