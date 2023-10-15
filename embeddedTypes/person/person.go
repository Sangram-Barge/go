package person

import "fmt"

type Person struct {
  FirstName string
  LastName string
  Age uint8
}

func (p *Person) SayHello() {
  fmt.Println("Hello from", p.FirstName, p.LastName )
}
