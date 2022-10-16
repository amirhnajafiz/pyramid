package pyramid

import (
	"container/heap"
	"fmt"
)

func main() {
	var pq Heap

	heap.Init(&pq)

	for i := 10; i > 0; i-- {
		heap.Push(&pq, &Item{value: i + 1, index: i})
	}

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d ", item.value)
	}
}
