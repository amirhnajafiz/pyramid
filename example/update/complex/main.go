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
	h := pyramid.NewHeap[*Data](func(a *Data, b *Data) bool {
		return a.Priority < b.Priority
	})

	special := &Data{
		Priority: 100,
		Data:     "i am special",
	}

	h.Push(special)

	for i := 2; i < 10; i++ {
		h.Push(&Data{Priority: i, Data: fmt.Sprintf("data: %d", i+1)})
	}

	special.Priority = 0

	h.Update(special, special, func(a *Data, b *Data) bool {
		return a.Data == b.Data
	})

	for h.Length() > 0 {
		fmt.Printf("%s\n", h.Pop().(*Data).Data)
	}
}
