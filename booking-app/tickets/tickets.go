package tickets

import (
  "booking-app/user"
)

func BookTicket(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8, bookings []user.UserData) (uint8, []user.UserData) {

  remainingTickets -= userTickets
  userData := user.UserData {
    FirstName: firstName,
    LastName: lastName,
    Email: email,
    NumberOfTickets: userTickets ,
  }

  return remainingTickets, append(bookings, userData)
}
