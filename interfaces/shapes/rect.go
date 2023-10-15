package shapes

type Rectangle struct {
  Length, Width float64
}

func (r *Rectangle) Area() float64 {
  return r.Length * r.Width
}
