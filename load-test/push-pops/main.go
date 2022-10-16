package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/amirhnajafiz/pyramid"
)

func main() {
	var (
		flNumberOfPush = flag.Int("push", 1000, "Number of push to heap")
		flNumberOfPop  = flag.Int("pop", 1000, "Number of pops from heap")
	)

	flag.Parse()

	h := pyramid.NewHeap[int](func(a any, b any) bool {
		return a.(int) < b.(int)
	})

	fmt.Printf("testing: %d numbers\n", *flNumberOfPop+*flNumberOfPush)
	fmt.Printf("start: %s\n", time.Now().Format(time.StampMilli))

	for *flNumberOfPop+*flNumberOfPush > 0 {
		num := rand.Int() % 100
		if num > 20 {
			h.Push(rand.Int())
			*flNumberOfPush--
		} else {
			if h.Length() > 0 {
				h.Pop()
				*flNumberOfPop--
			}
		}
	}

	fmt.Printf("done: %s\n", time.Now().Format(time.StampMilli))
	fmt.Printf("final size: %d\n", h.Length())
}
