package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	for {

		firstName, lastName, userEmail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, userEmail)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, userEmail)

			firstNames := getFirstNames()
			fmt.Printf("Bookings so far: %v\n", firstNames)

			if remainingTickets == 0 {
				// End program
				fmt.Println("Our Conference is booked out. Come back next year")
				break
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of : %v tickets and remaining tickets are: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets // update remaining tickets

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           userEmail,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("Remaining tickets are %v for %v\n", remainingTickets, conferenceName)

	//call function to get first names of bookings
}

func sendTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket: %v to email address %v\n", tickets, userEmail)
	fmt.Println("####################")
	wg.Done()
}
