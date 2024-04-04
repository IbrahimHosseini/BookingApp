package main

import (
	"fmt"
)

func main() {
	// Welcome message
	const companyName = "Flight Booking"
	fmt.Printf("Welecom to %v.\n", companyName)

	totalTickets := 100
	var remainingTickets = totalTickets

	var name string
	var family string
	var ticket int

	fmt.Println("Enter name: ")
	fmt.Scan(&name)

	fmt.Println("Enter name: ")
	fmt.Scan(&family)

	fmt.Println("How many ticket(s) do you want to buy?")
	fmt.Scan(&ticket)

	remainingTickets -= ticket

	fmt.Printf("Thanks, %v %v you buy %v ticke(s).\n", name, family, ticket)
	fmt.Printf("The %v ticket(s) is remaining.\n", remainingTickets)

}
