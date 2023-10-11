package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
  const conferenceTickets uint8 = 50
  var remainingTickets uint8 = conferenceTickets

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
  fmt.Printf("We have total of %v tickets and %v tickets are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

  var firstName string
  var lastName string
  var email string
  var userTickets uint8

  fmt.Println("Hello user enter your first name")
  fmt.Scan(&firstName)
  fmt.Println("Hello user enter your last name")
  fmt.Scan(&lastName)
  fmt.Println("Hello user enter your email")
  fmt.Scan(&email)
  fmt.Printf("Hello %v %v enter how many tickets you want\n", firstName, lastName)
  fmt.Scan(&userTickets)

  remainingTickets -= userTickets

  fmt.Printf("Thank you %v %v for booking %v tickets. Ticket will be delivered at %v address\n", firstName, lastName, userTickets, email)
}
