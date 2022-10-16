<h1 align="center">
Pyramid
</h1>

Fast generic type **min-heap** in Golang.

## How to use?

```shell
go get github.com/amirhnajafiz/pyramid
```

## Example

Creating a max-heap of integers type.

```go
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
```

Creating a min-heap of custome data type.

```go
type Data struct {
	Priority int
	Data     string
}

func main() {
	h := pyramid.NewHeap[Data](func(a any, b any) bool {
		return a.(Data).Priority < b.(Data).Priority
	})

	for i := 0; i < 10; i++ {
		h.Push(Data{Priority: i, Data: fmt.Sprintf("data: %d", i+1)})
	}

	for h.Length() > 0 {
		fmt.Printf("%s\n", h.Pop().(Data).Data)
	}
}
```
