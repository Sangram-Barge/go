package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
  file, err := os.Open("test.txt")
  if err != nil {
    panic(fmt.Sprintf("some error %v\n", err))
  }
  defer file.Close()

  stat, err := file.Stat()
  if err != nil {
    return
  }

  fmt.Println("statistics of file", stat)

  bs := make([]byte, stat.Size())
  _, err = file.Read(bs)

  if err != nil {
    return
  }

  fmt.Println(strings.Repeat("_", 20))
  fmt.Println(bs)
  fmt.Println(strings.Repeat("_", 20))
  fmt.Println(string(bs))
  fmt.Println(strings.Repeat("_", 20))

  byteStream := []byte("hello world this is sangram barge")
  msg, err := file.Write(byteStream)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("writter bytes", msg) 
}
