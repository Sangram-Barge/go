package main

import (
  "fmt"
  "strings"
)

func main() {
  line()
  var conferenceName string = "Go Conference"
  const conferenceTickets uint8 = 50
  var remainingTickets uint8 = conferenceTickets

  greet(conferenceName, conferenceTickets, remainingTickets)

  bookings := []string{}

  for {
    firstName, lastName, isValid := getUserName()
    if !(isValid) {
      fmt.Printf("user name invalid, firstname and last name should be more than two characters\n")
      line()
      continue
    }

    email, isValid := getUserEmail(firstName, lastName)
    if !(isValid) {
      fmt.Printf("user email invalid must contain '@' and '.com'\n")
      line()
      continue
    } 

    userTickets, isValid := getUserTickets(firstName, lastName)
    if !(isValid) {
      fmt.Printf("Cannot book more than remaining Tickets, Remaining Tickets => %v\n", remainingTickets)
      line()
      continue
    }

    remainingTickets -= userTickets
    bookings = append(bookings, firstName + " " + lastName)

    provideConfirmation(firstName, lastName, email, userTickets, remainingTickets, conferenceName)

    firstNames := getFirstNames(bookings)
    fmt.Printf("Theses are all the bookings in the app %v \n", firstNames)
    line()

    if remainingTickets == 0 {
      fmt.Println("We are full all tickets are booked, please come again next year")
      break
    }
  }
}

func greet(confName string, confTickets uint8, remainingTickets uint8) {
  fmt.Printf("Welcome to %v booking app\n", confName)
  fmt.Printf("We have total of %v tickets and %v tickets are available\n", confTickets, remainingTickets)
  fmt.Println("Get your tickets here to attend")
}

func line() {
  fmt.Println("--------------------------------------------------------------------------")
}

func getUserTickets(firstName string, lastName string) (userTickets uint8, err bool) {
    fmt.Printf("Hello %v %v enter how many tickets you want\n", firstName, lastName)
    fmt.Scan(&userTickets)
    return userTickets, (userTickets > 0 && userTickets <= userTickets)
}

func getFirstNames(names []string) (firstNames []string) {
  for _, name := range names {
    firstNames = append(firstNames, strings.Fields(name)[0])
  }
  return firstNames
}

func getUserName() (firstName string, lastName string, err bool) {
  fmt.Println("Hello user enter your first name")
  fmt.Scan(&firstName)
  fmt.Printf("Hello %v enter your last name\n", firstName)
  fmt.Scan(&lastName)
  return firstName, lastName, (len(firstName) >= 2 && len(lastName) >= 2)
}

func getUserEmail(firstName string, lastName string) (email string, err bool) {
  fmt.Printf("Hello %v enter your email\n", firstName + " " + lastName)
  fmt.Scan(&email)
  return email, (strings.Contains(email, "@") && strings.Contains(email, ".com"))
}

func provideConfirmation(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8, conferenceName string) {
  fmt.Printf("Thank you %v for booking %v tickets. Ticket will be delivered at %v address\n", firstName + " " + lastName, userTickets, email)
  fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}








