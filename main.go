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
	var ticket int

	buyTicket(name, family, ticket)
}

// buy ticket process
func buyTicket(name string, family string, ticket int) {

	var remainingTickets = totalTickets

	fmt.Printf("\nThere are %v ticket(s) available.\n", remainingTickets)

	for remainingTickets > 0 {
		fmt.Println("\nEnter your firstname: ")
		fmt.Scan(&name)

		fmt.Println("Enter your lastname: ")
		fmt.Scan(&family)

		fmt.Println("How many ticket(s) do you want to buy?")
		fmt.Scan(&ticket)

		if remainingTickets >= ticket {
			remainingTickets -= ticket

			fmt.Printf("Thanks, %v %v buy %v ticke(s).\n", name, family, ticket)
			fmt.Printf("The %v ticket(s) is remaining.\n", remainingTickets)
		} else {
			fmt.Printf("Your request is more than remainig. The remaining tickets are %v.\n", remainingTickets)
		}

	}

	fmt.Printf("The tickets sold out.\n")
}
