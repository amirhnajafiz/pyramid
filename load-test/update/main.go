package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/amirhnajafiz/pyramid"
)

type Data struct {
	id    int
	value int
}

func main() {
	var (
		flNumberOfPush    = flag.Int("push", 1000, "Number of push to heap")
		flNumberOfUpdates = flag.Int("update", 1000, "Number of updates to heap")
	)

	flag.Parse()

	var items []*Data

	h := pyramid.NewHeap[*Data](func(a *Data, b *Data) bool {
		return a.value < b.value
	})

	fmt.Printf("testing: %d numbers\n", *flNumberOfPush)
	fmt.Printf("start: %s\n", time.Now().Format(time.StampMilli))

	for i := 0; i < *flNumberOfPush; i++ {
		d := &Data{
			id:    i,
			value: 1 + rand.Int()%100000,
		}

		h.Push(d)
		items = append(items, d)
	}

	for i := 0; i < *flNumberOfUpdates; i++ {
		index := rand.Int() % len(items)
		newItem := &Data{
			id:    index,
			value: rand.Int() % 100000,
		}

		h.Update(items[index], newItem, func(a any, b any) bool {
			return a.(*Data).id == b.(*Data).id
		})
	}

	fmt.Printf("done: %s\n", time.Now().Format(time.StampMilli))
	fmt.Printf("final size: %d\n", h.Length())
	fmt.Printf("min: %d\n", h.Pop().(*Data).value)
}
