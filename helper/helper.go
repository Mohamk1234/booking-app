package helper

import "strings"

func Validate(firstName string, lastName string, userTickets uint, remainingTickets uint, email string) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets <= remainingTickets && userTickets > 0

	return isValidName, isValidEmail, isValidTickets
}
