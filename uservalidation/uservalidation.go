package uservalidation

func ValidateUserInput(firstName, lastName string, userTickets, remainingTickets uint) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 1
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidTicketNumber
}
