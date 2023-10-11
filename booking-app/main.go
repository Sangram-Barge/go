package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println("--------------------------------------------------------------------------")
	var conferenceName string = "Go Conference"
  const conferenceTickets uint8 = 50
  var remainingTickets uint8 = conferenceTickets

  bookings := []string{}

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
  fmt.Printf("We have total of %v tickets and %v tickets are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

  for {
    var firstName string
    var lastName string
    var email string
    var userTickets uint8

    fmt.Println("Hello user enter your first name")
    fmt.Scan(&firstName)
    fmt.Printf("Hello %v enter your last name\n", firstName)
    fmt.Scan(&lastName)
    fmt.Printf("Hello %v enter your email\n", firstName + " " + lastName)
    fmt.Scan(&email)
    fmt.Printf("Hello %v %v enter how many tickets you want\n", firstName, lastName)
    fmt.Scan(&userTickets)

    remainingTickets -= userTickets
    bookings = append(bookings, firstName + " " + lastName)

    fmt.Printf("Thank you %v for booking %v tickets. Ticket will be delivered at %v address\n", bookings[0], userTickets, email)
    fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
    
    firstNames := []string{}
    for _, booking := range bookings {
      firstNames = append(firstNames, strings.Fields(booking)[0])
    }
    fmt.Printf("Theses are all the bookings in the app %v \n", firstNames)
  }
}