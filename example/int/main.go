package main

import (
	"fmt"

	"github.com/amirhnajafiz/pyramid"
)

func main() {
	h := pyramid.NewHeap[int](func(a any, b any) bool {
		return a.(int) > b.(int)
	})

	h.Push(2)
	h.Push(12)
	h.Push(4)
	h.Push(90)
	h.Push(20)

	for h.Length() > 0 {
		fmt.Printf("%d ", h.Pop().(int))
	}
}
