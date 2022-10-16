package pyramid

// An item is something we manage in a priority queue.
type item[T any] struct {
	// each item has a value of its own that could be anything.
	value T
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// compareFunction
// is used to compare items in priority queue.
type compareFunction func(any, any) bool

// equalFunction
// is used to update elements in priority queue.
type equalFunction func(any, any) bool
