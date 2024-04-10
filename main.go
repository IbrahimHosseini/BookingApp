package main

import (
	"fmt"
)

const totalTickets = 100

func main() {
	// Welcome message
	const companyName = "Flight Booking"
	fmt.Printf("Welecom to %v.\n", companyName)

	var name string
	var family string
	var email string
	var ticket int

	buyTicket(name, family, email, ticket)
}

// buy ticket process
func buyTicket(name string, family string, email string, ticket int) {

	var remainingTickets = totalTickets

	fmt.Printf("\nThere are %v ticket(s) available.\n", remainingTickets)

	for remainingTickets > 0 {
		fmt.Println("\nEnter your firstname: ")
		fmt.Scan(&name)

		fmt.Println("Enter your lastname: ")
		fmt.Scan(&family)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("How many ticket(s) do you want to buy?")
		fmt.Scan(&ticket)

		isValidName, isValidEmail, isValidTicketNumber := inputValidation(name, family, remainingTickets, ticket, email)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets -= ticket

			fmt.Printf("Thanks, %v %v buy %v ticke(s).\nTicket(s) will be sent to %v.", name, family, ticket, email)
			fmt.Printf("The %v ticket(s) is remaining.\n", remainingTickets)
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is not valid.")
			}
		}

	}

	fmt.Printf("The tickets sold out.\n")
}
