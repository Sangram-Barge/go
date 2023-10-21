package main

import (
	"fmt"
	"sort"
)

type Person struct {
  Name string
  Age uint8
}

type ByName []Person

func (bn ByName) Len() int {
  return len(bn)
}

func (bn ByName) Less(i, j int) bool {
  return bn[i].Age < bn[j].Age
}

func (bn ByName) Swap(i, j int) {
  bn[i], bn[j] = bn[j], bn[i]
}

func main() {
  kids := []Person{
    {
      Name: "sangram",
      Age: 12,
    },
    {
      Name: "anushree",
      Age: 11,
    },
  }

  sort.Sort(ByName(kids))
  fmt.Println(kids)
}
