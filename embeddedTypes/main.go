package main

import "embeddedTypes/person"

func main() {
  var p = person.Person{
    FirstName: "Sangram",
    LastName: "Barge",
    Age: 25,
  }

  var a = person.Android{
    Model: "T800",
    Person: person.Person{
    FirstName: "Arnold",
    LastName: "Schvez",
    Age: 89,
    },
  }

  p.SayHello()
  a.SayHello()

}
