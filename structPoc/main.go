package main

import (
	"fmt"
	"math"
)

type Circle struct {
  x, y, r float32
}

func main() {
  var c Circle

  fmt.Println(c)

  var d = new(Circle)

  fmt.Println(d)

  s := Circle{
    0, 0, 2,
  }
  fmt.Println(s)

  s2 := &Circle{
    0, 0, 4,
  }
  fmt.Println(s2)
  
  CircleArea(&c)
  CircleArea(d)
  CircleArea(&s)
  CircleArea(s2)
}

func CircleArea(c *Circle) {
  fmt.Printf("area fo circle %v is %v \n", c, math.Pi * c.r * c.r)
}
