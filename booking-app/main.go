package main

import (
	"fmt"
	"go-learning/booking-app/helper"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTickets(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of our bookings are: %v\n", firstNames)
			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
			}
		} else {
			if !isValidName {
				fmt.Println("Your firstname and/or lastname is too short")
			}
			if !isValidEmail {
				fmt.Println("Your Email is invalid")
			}
			if !isValidUserTickets {
				fmt.Println("Your ticket number is not valid")
			}
		}
	}

	city := "London"

	switch city {
	case "New York":
		//execute code
	case "Singapure", "Hong Kong":
		//execute code
	case "London", "Berlin":
		//execute code
	case "Mexico City":
		//execute code
	default:
		fmt.Println("No valid city selected")
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, bookings := range bookings {
		var names = strings.Fields(bookings)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number your tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidUserTickets
}
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
