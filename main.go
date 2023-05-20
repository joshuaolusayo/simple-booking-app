package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets

// var bookings [50]string // array
// var bookings []string // slice
// var bookings = make([]map[string]string, 0) // slice alternative declaration for the method above
var bookings = make([]UserData, 0) // slice alternative declaration for the method above

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetings()

	// for {
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if !(isValidName && isValidEmail && isValidTicket) {
		fmt.Println("\nInvalid data. Please, check your input and try again")
		fmt.Println("Your first name and last name must have length > 2")
		fmt.Printf("The number of ticket you want to book must not exceed %v\n\n", remainingTickets)
		// continue
	}

	bookTickets(userTickets, firstName, lastName, email)

	wg.Add(1)
	go sendTickets(userTickets, firstName, lastName, email)

	firstNames := getFirstNames()

	fmt.Printf("The first names of bookings are: %v\n", firstNames)

	if remainingTickets == 0 {
		// end program
		fmt.Println("Our conference is booked out. Check back next year")
		// break
	}

	// }

	wg.Wait()
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets do you want to book? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func greetings() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total tickets: %v Remaining tickets: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to apply")
	fmt.Printf("Conference name is %T; Conference ticket is %T; Remaining ticket is %T\n", conferenceName, conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// var firstName = names[0] // similar to aarray.split("") in javascript
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// bookings[0] = firstName + " " + lastName // array implementation
	bookings = append(bookings, userData)

	fmt.Printf("%v %v booked %v tickets. Ticket details will be sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remain for %v event\n", remainingTickets, conferenceName)

}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n#################")
	fmt.Printf("Sending tiket:\n %v \n to \n %v address\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
