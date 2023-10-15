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

  circle := Circle{
    2, 3, 9,
  }
  circle.Area()
  
  var r = Rectangle{
    3, 4,
  }

  r.Area()
}

func CircleArea(c *Circle) {
  fmt.Printf("this a function to which we are passing circle area fo circle %v is %v \n", c, math.Pi * c.r * c.r)
}

func (c *Circle) Area() {
  fmt.Printf("this is a method of circle area fo circle %v is %v \n", c, math.Pi * c.r * c.r)
}

type Rectangle struct {
  l, w float32
}

func (r *Rectangle) Area() {
  fmt.Println("this is method of Rectangle, area of rect is", r.l * r.w)
}





