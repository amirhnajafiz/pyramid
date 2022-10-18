package main

import (
	"fmt"

	"github.com/amirhnajafiz/pyramid"
)

type Data struct {
	Priority int
	Data     string
}

func main() {
	h := pyramid.NewHeap[Data](func(a Data, b Data) bool {
		return a.Priority < b.Priority
	})

	for i := 0; i < 10; i++ {
		h.Push(Data{Priority: i, Data: fmt.Sprintf("data: %d", i+1)})
	}

	for h.Length() > 0 {
		fmt.Printf("%s\n", h.Pop().(Data).Data)
	}
}
