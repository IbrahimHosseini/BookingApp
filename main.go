package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const totalTickets = 100

var remainingTickets = totalTickets
var ticket int

const confranceName = "Flight Booking"

// create a list of maps
var bookings = make([]UserData, 0)

// struct
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {
	// Welcome message
	fmt.Printf("Welecom to %v.\n", confranceName)

	var name string
	var family string
	var email string

	buyTicket(name, family, email)
}

// buy ticket process
func buyTicket(name string, family string, email string) {

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

		isValidName, isValidEmail, isValidTicketNumber := helper.InputValidation(name, family, remainingTickets, ticket, email)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(ticket, name, family, email)

			// user the `go` keyword to useing cuncurency in app and run the thread in background.
			go sendTicket(ticket, name, family, email)

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

func bookTicket(userTickets int, firstName string, lastName string, email string) {

	remainingTickets -= ticket

	// creat a map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is:\n %v\n", bookings)

	fmt.Printf("Thanks, %v %v buy %v ticke(s).\nTicket(s) will be sent to %v.", firstName, lastName, userTickets, email)
	fmt.Printf("The %v ticket(s) is remaining for %v.\n", remainingTickets, confranceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf("%v ticket(s) for %v %v.", userTickets, firstName, lastName)

	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v.\n", ticket, email)
	fmt.Println("########################")
}
