<p align="center">
<img src="assets/logo.webp" width="700"  alt="logo"/>
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

Fast generic **Heap** in Golang. Create your Heap data structure with any data type. Set your
compression function to sort your data with any field that you want. Push, Pop and Update your
data in heap without any problems. Write your equality function to filter
the elements that you want to update.

**Pyramid** is built with **Golang** internal libraries and does not include any external
libraries.

## How to use?

Get **pyramid** package.

```shell
go get github.com/amirhnajafiz/pyramid
```

Now you can use pyramid to build any type of **Heap**.

## Example

Creating a _max-heap_ of type Integer.

```go
package main

import (
	"fmt"

	"github.com/amirhnajafiz/pyramid"
)

func main() {
	// initializing a heap with comparator function
	h := pyramid.NewHeap[int](func(a int, b int) bool {
		return a > b
	})

	// Using push method
	h.Push(2)
	h.Push(12)
	h.Push(4)
	h.Push(90)
	h.Push(20)

	// Using length method
	for h.Length() > 0 {
		// Using pop method
		fmt.Printf("%d ", h.Pop().(int))
	}
}
```

Creating a _min-heap_ of type custom _Data_.

```go
// Creating a custom Data struct type.
type Data struct {
	Priority int
	Data     string
}

func main() {
	// Creating a heap of Data type.
	h := pyramid.NewHeap[Data](func(a Data, b Data) bool {
		return a.Priority < b.Priority
	})

	// Push data into heap.
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
