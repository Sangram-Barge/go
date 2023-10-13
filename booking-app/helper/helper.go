package helper

import "fmt"

func Greet(confName string, confTickets uint8, remainingTickets uint8) {
  fmt.Printf("Welcome to %v booking app\n", confName)
  fmt.Printf("We have total of %v tickets and %v tickets are available\n", confTickets, remainingTickets)
  fmt.Println("Get your tickets here to attend")
}

func Line() {
  fmt.Println("--------------------------------------------------------------------------")
}

func ProvideConfirmation(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8, conferenceName string) {
  fmt.Printf("Thank you %v for booking %v tickets. Ticket will be delivered at %v address\n", firstName + " " + lastName, userTickets, email)
  fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
