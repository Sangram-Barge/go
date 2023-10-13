package main

import (
  "fmt"
  "booking-app/helper"
  "booking-app/user"
)

func main() {
  helper.Line()
  var conferenceName string = "Go Conference"
  const conferenceTickets uint8 = 50
  var remainingTickets uint8 = conferenceTickets

  helper.Greet(conferenceName, conferenceTickets, remainingTickets)

  bookings := []string{}

  for {
    firstName, lastName, isValid := user.GetUserName() 
    if !(isValid) {
      fmt.Printf("user name invalid, firstname and last name should be more than two characters\n")
      helper.Line()
      continue
    }

    email, isValid := user.GetUserEmail(firstName, lastName)
    if !(isValid) {
      fmt.Printf("user email invalid must contain '@' and '.com'\n")
      helper.Line()
      continue
    } 

    userTickets, isValid := user.GetUserTickets(firstName, lastName, remainingTickets)
    if !(isValid) {
      fmt.Printf("Cannot book more than remaining Tickets, Remaining Tickets => %v\n", remainingTickets)
      helper.Line()
      continue
    }

    remainingTickets -= userTickets
    bookings = append(bookings, firstName + " " + lastName)

    helper.ProvideConfirmation(firstName, lastName, email, userTickets, remainingTickets, conferenceName)

    firstNames := user.GetFirstNames(bookings)
    fmt.Printf("Theses are all the bookings in the app %v \n", firstNames)
    helper.Line()

    if remainingTickets == 0 {
      fmt.Println("We are full all tickets are booked, please come again next year")
      break
    }
  }
}
