package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "GO conference"

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
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.Validate(firstName, lastName, userTickets, remainingTickets, email)

		if isValidEmail && isValidName && isValidTickets {
			remainingTickets = remainingTickets - uint(userTickets)

			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: userTickets,
			}

			bookings = append(bookings, userData)

			fmt.Printf("Thank you %v for booking %v tickets, you will receive a confirmation mail shortly on %v\n", firstName, userTickets, email)
			fmt.Printf("There are %v tickets remaining\n", remainingTickets)

			printFirstNames()

			wg.Add(1)
			go sendTicket(firstName, lastName, userTickets, email)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is completely booked\n")
				break
			}
		} else {

			if !isValidName {
				fmt.Printf("first or lastname is too short\n")
			}

			if !isValidEmail {
				fmt.Printf("Email address is not valid\n")
			}

			if !isValidTickets {
				fmt.Printf("We only have %v tickets remaining, please enter a valid ticket number\n", remainingTickets)
			}

			fmt.Printf("Please try Booking again\n")

		}

	}

	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets remaining\n", remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() {
	firstNames := []string{}
	for _, value := range bookings {
		firstNames = append(firstNames, value.firstName)
	}
	fmt.Printf("This is the complete booking list %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter the amount of tickets you want to book")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func sendTicket(firstName string, lastName string, userTickets uint, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	time.Sleep(10 * time.Second)
	fmt.Printf("##################\n")
	fmt.Printf("Sending ticket:\n %v \n to email address %v", ticket, email)
	fmt.Printf("##################\n")
	wg.Done()
}
