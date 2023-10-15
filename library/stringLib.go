package main

import (
	"fmt"
	"strings"
)

func StringTest() {
  fmt.Println("Sangram contains san?",strings.Contains("sangram", "san"))
  fmt.Println("Occurence of 'a' in sangrambarge", strings.Count("sangrambarge", "a"))
  fmt.Println("Does Sangram start with gram?", strings.HasPrefix("Sangram", "gram"))
  fmt.Println("Does Sangram end with gram?", strings.HasSuffix("Sangram", "gram"))
  fmt.Println("Index of g and e in sangram", strings.Index("sangram", "g"), strings.Index("sangram", "e"))
  fmt.Println(strings.Repeat("-", 20))
  myStrings := []string{
    "Hello",
    "World",
    "This is",
    "a slice of string",
    "which would be converted into single string",
  }
  fmt.Println("concatinated string with semicolon", strings.Join(myStrings, ";"))
  fmt.Println("replce a with u from sangram", strings.Replace("sangram", "a", "u", -1))
  fmt.Println(strings.Repeat("-", 20))
  fmt.Println("split this-is-my-world-into saperate strings", strings.Split("this-is-my-world", "-"))

}
