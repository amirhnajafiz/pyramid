package test

import (
	"github.com/amirhnajafiz/pyramid"
	"testing"
)

func TestHeap(t *testing.T) {
	h := pyramid.NewHeap[int64](func(a any, b any) bool {
		return a.(int64) < b.(int64)
	})

	h.Push(int64(100))
	h.Push(int64(1))
	h.Push(int64(10))

	if h.Length() != 3 {
		t.Error("push failed")
	}

	if h.Pop().(int64) != 1 {
		t.Error("heap structure failed")
	}
}

func TestInvalidType(t *testing.T) {
	type item struct {
		value int
	}

	h := pyramid.NewHeap[item](func(a any, b any) bool {
		return a.(item).value < b.(item).value
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
