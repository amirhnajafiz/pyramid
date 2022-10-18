<p align="center">
<img src="assets/logo.webp" width="700" />
</p>

<p align="center">
<img src="https://img.shields.io/badge/Golang-1.19-66ADD8?style=for-the-badge&logo=go" alt="go version" />
<img src="https://img.shields.io/badge/Version-0.1.3-DD1199?style=for-the-badge&logo=github" alt="version" />
<img src="https://img.shields.io/badge/Load_Test-1M-442266?style=for-the-badge&logo=k6" alt="version" />
<br />
</p>

<h1 align="center">
Pyramid
</h1>

Fast generic **Heap** in Golang. Create your heap with any type of data. Set your
compression function to sort your data with any field that you want. Push, Pop and Update your
data in heap without any problems.

## How to use?

Get **pyramid** package.

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
```

Creating a min-heap of custome data type.

```go
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
```

### Updating a reference

You can update any item in the list, also you can define a 
equal function to check the equality of your objects.

```go
special := 675

h.Push(special)

for i := 2; i < 100; i++ {
    h.Push(i)
}

h.Update(special, 0, func(a int, b int) bool {
    return a == b
})
```

## Load test
If we have 10000 items, and we want to update them 1 Million times, it would
only take 12 seconds:

```shell
go run load-test/update/main.go -push 10000 -update 1000000
```

```shell
testing: 10000 numbers
start: Oct 16 14:42:34.752
done: Oct 16 14:42:52.064
final size: 10000
min: 2
```
