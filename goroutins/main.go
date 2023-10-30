package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func main() {
  for i := 0; i < 20; i++ {
    go calculate(i)
  }
  var s string
  fmt.Scanln(&s)
}

func calculate(i int) {
  fmt.Println(i)
  //amt := time.Duration(rand.Intn(250))
  //time.Sleep(time.Millisecond * amt)
}
