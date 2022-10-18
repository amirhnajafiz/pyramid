package main

import (
	"fmt"

	"github.com/amirhnajafiz/pyramid"
)

func main() {
	h := pyramid.NewHeap[int](func(a int, b int) bool {
		return a < b
	})

	special := 675

	h.Push(special)

	for i := 2; i < 100; i++ {
		h.Push(i)
	}

	h.Update(special, 0, func(a any, b any) bool {
		return a.(int) == b.(int)
	})

	for h.Length() > 0 {
		fmt.Printf("%d\n", h.Pop().(int))
	}
}
