<p align="center">
<img src="assets/logo.webp" />
</p>

<p align="center">
<img src="https://img.shields.io/badge/Golang-1.19-66ADD8?style=for-the-badge&logo=go" alt="go version" />
<img src="https://img.shields.io/badge/Version-0.1.1-purple?style=for-the-badge&logo=github" alt="version" /><br />
</p>

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
