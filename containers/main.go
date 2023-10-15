package main

import (
	"container/list"
	"fmt"
)

func main() {
  var x list.List
  x.PushFront(3)
  x.PushFront("hello")
  x.PushFront(3.4)

  for e:= x.Front(); e != nil; e = e.Next() {
    fmt.Printf("value %v\n", e.Value)
  }
}
