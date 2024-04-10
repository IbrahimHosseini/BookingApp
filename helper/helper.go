package helper

import "strings"

func InputValidation(name string, family string, remainingTickets int, ticket int, email string) (bool, bool, bool) {

	isValidName := (len(name) >= 2 && len(family) >= 2)
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := remainingTickets >= ticket && ticket > 0

	return isValidName, isValidEmail, isValidTicketNumber
}
