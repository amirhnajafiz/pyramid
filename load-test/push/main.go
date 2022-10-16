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
	)

	flag.Parse()

	h := pyramid.NewHeap[int](func(a any, b any) bool {
		return a.(int) < b.(int)
	})

	fmt.Printf("testing: %d numbers\n", *flNumberOfPush)
	fmt.Printf("start: %s\n", time.Now().Format(time.StampMilli))

	for i := 0; i < *flNumberOfPush; i++ {
		h.Push(rand.Int() % 100)
	}

	fmt.Printf("done: %s\n", time.Now().Format(time.StampMilli))
	fmt.Printf("final size: %d\n", h.Length())
}
