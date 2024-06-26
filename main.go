package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const totalTickets = 100

var remainingTickets = totalTickets
var ticket int

const confranceName = "Flight Booking"

// create a list of maps
var bookings = make([]UserData, 0)

// creat a wait group
var wg = sync.WaitGroup{}

// The UserData struct defines the structure of user data including first name, last name, email, and
// number of tickets.
// @property {string} firstName - The `firstName` property in the `UserData` struct represents the
// first name of a user. It is of type `string` and stores the first name of the user.
// @property {string} lastName - The `lastName` property in the `UserData` struct represents the last
// name of a user. It is of type `string` and is used to store the last name of the user.
// @property {string} email - The `email` property in the `UserData` struct represents the email
// address of the user.
// @property {int} numberOfTickets - The `numberOfTickets` property in the `UserData` struct represents
// the number of tickets associated with a user. It is an integer type that stores the count of tickets
// owned by the user.
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

	// for remainingTickets > 0 {
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

		// agg wait group for function
		wg.Add(1)
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

	// }

	fmt.Printf("The tickets sold out.\n")

	wg.Wait()
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

	wg.Done()
}
