package main

import(
  "fmt"
  "interfaces/shapes"
)

func main() {
  s := shapes.Circle{
    Radius: 3.4,
  }

  s2 := shapes.Circle{
    Radius: 4.2,
  }

  rect1 := shapes.Rectangle{
    Length: 9.8, 
    Width: 8.6,
  }

  totalArea := calculateTotalArea(&s, &s2, &rect1)

  fmt.Println("total area is", totalArea)


  mulShape := shapes.MultiShape{
    Shapes: []shapes.Shape{
      &shapes.Circle{Radius: 3},
    },
  }

  fmt.Println("multishape implementation total area", mulShape.Area())

}

func calculateTotalArea(shapess ...shapes.Shape) (total float64){
  for _, s := range shapess {
    total += s.Area()
  }
  return total
}
