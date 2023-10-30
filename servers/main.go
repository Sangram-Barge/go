package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
  conn, err := net.Listen("tcp", ":9999")
  if err != nil {
    fmt.Println(err)
    return
  }

  for {
    c, err := conn.Accept()
    if err != nil {
      fmt.Println(err)
      continue
    }
    go handleServerConnection(c)
  }
}

func handleServerConnection(c net.Conn) {
  defer c.Close()
  var msg string
  err := gob.NewDecoder(c).Decode(&msg)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("received", msg)
  }
}


func main() {
  go server()

  var t string
  fmt.Scanln(&t)
}











