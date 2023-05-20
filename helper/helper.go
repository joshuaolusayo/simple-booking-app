package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicket
}