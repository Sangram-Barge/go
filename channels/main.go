package main

import (
	"fmt"
	"time"
)

func main() {
  var c chan string = make(chan string)
  go pinger(c) 
  go printer(c) 
  go ponger(c)
  var t string
  fmt.Scanln(&t)
}

func pinger (c chan<- string) {
  for {
    c <- "ping"
  }
}

func ponger (c chan<- string) {
  for {
    c <- "pong"
  }
}

func printer (c <-chan string) {
  for {
    fmt.Println(<-c)
    time.Sleep(time.Second * 2)
  }
}

