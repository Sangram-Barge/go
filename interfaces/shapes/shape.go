package shapes

type Shape interface {
  Area() float64;
}

type MultiShape struct {
  Shapes []Shape
}


func (sh *MultiShape) Area() (total float64) {
  for _, s := range sh.Shapes {
    total += s.Area()
  }
  return total
}
