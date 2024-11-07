package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()
	for remainingTickets > 0 && len(bookings) < 50 {

		// ask user for their name
		firstName, lastName, userEmail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userEmail, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, bookings, firstName, lastName, userEmail)

			firstNames := getFirstNames()
			fmt.Println("Bookings so far: ", firstNames)

			if remainingTickets == 0 {
				// End program
				fmt.Println("Our Conference is booked out. Come back next year")
				break
			} else {
				fmt.Printf("Sorry we only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email does not contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of : %v tickets and remaining tickets are: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("conferenceTickets: is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, conferenceName)

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, userEmail string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&userEmail)

	fmt.Println("How many tickets you want to book?")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(userTickets uint, bookings []string, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets // update remaining tickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("Remaining tickets are %v for %v\n", remainingTickets, conferenceName)

	//call function to get first names of bookings

}
