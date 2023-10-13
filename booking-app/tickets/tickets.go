package tickets

func BookTicket(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8, bookings []map[string]any) (uint8, []map[string]any) {

  remainingTickets -= userTickets
  user := make(map[string]any)

  user["firstName"] = firstName
  user["lastName"] = lastName
  user["email"] = email
  user["userTickets"] = userTickets

  return remainingTickets, append(bookings, user)
}
