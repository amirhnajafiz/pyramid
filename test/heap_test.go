package test

import (
	"testing"

	"github.com/amirhnajafiz/pyramid"
)

// TestHeap
// testing heap data structure with int64 type.
func TestHeap(t *testing.T) {
	// creating a new heap
	h := pyramid.NewHeap[int64](func(a int64, b int64) bool {
		return a < b
	})

	// push some data into heap
	h.Push(int64(100))
	h.Push(int64(1))
	h.Push(int64(10))

	// check push result
	if h.Length() != 3 {
		t.Error("push failed")
	}

	// check pop operation
	if h.Pop().(int64) != 1 {
		t.Error("heap structure failed")
	}
}

// TestValidType
// testing heap of a valid custom type.
func TestValidType(t *testing.T) {
	type item struct {
		value int
	}

	h := pyramid.NewHeap[item](func(a item, b item) bool {
		return a.value < b.value
	})

	h.Push(item{
		2,
	})
	h.Push(item{
		1,
	})

	if h.Length() != 2 {
		t.Error("push failed")
	}

	if h.Pop().(item).value != 1 {
		t.Error("pop failed")
	}
}
