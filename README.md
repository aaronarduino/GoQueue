# goqueue
Queue data structure in Golang. For integers only.

## To use:

Just go get the package.
`go get github.com/aaronarduino/goqueue`

Then use.
```
package main

import (
  "fmt"
  "github.com/aaronarduino/goqueue"
)

func main() {
  // Create queue
  testq := goqueue.New()
  
  // Push some values
  testq.Push(1)
  testq.Push(2)
  
  // Pop some values
  fmt.Println("Popped:", testq.Pop)
  fmt.Println("Popped:", testq.Pop)
}
```
