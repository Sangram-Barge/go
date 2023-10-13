package user

import (
  "strings"
  "fmt"
)

func GetFirstNames(names []string) (firstNames []string) {
  for _, name := range names {
    firstNames = append(firstNames, strings.Fields(name)[0])
  }
  return firstNames
}

func GetUserName() (firstName string, lastName string, err bool) {
  fmt.Println("Hello user enter your first name")
  fmt.Scan(&firstName)
  fmt.Printf("Hello %v enter your last name\n", firstName)
  fmt.Scan(&lastName)
  return firstName, lastName, (len(firstName) >= 2 && len(lastName) >= 2)
}

func GetUserEmail(firstName string, lastName string) (email string, err bool) {
  fmt.Printf("Hello %v enter your email\n", firstName + " " + lastName)
  fmt.Scan(&email)
  return email, (strings.Contains(email, "@") && strings.Contains(email, ".com"))
}

func GetUserTickets(firstName string, lastName string, remainingTickets uint8) (userTickets uint8, err bool) {
    fmt.Printf("Hello %v %v enter how many tickets you want\n", firstName, lastName)
    fmt.Scan(&userTickets)
    return userTickets, (userTickets > 0 && userTickets <= remainingTickets)
}
