package main

import (
	"fmt"

	"github.com/amirhnajafiz/pyramid"
)

func main() {
	h := pyramid.NewHeap[int](func(a int, b int) bool {
		return a > b
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
